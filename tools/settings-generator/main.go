// Settings Code Generator - generates boilerplate code for settings management
// Usage: go run tools/settings-generator/main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// SettingsSchema defines the structure of settings schema
type SettingsSchema struct {
	Meta     Meta                  `json:"_meta"`
	Settings map[string]SettingDef `json:"settings"`
}

type Meta struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

type SettingDef struct {
	Type        string      `json:"type"` // int, string, bool
	Default     interface{} `json:"default"`
	Category    string      `json:"category"`
	Encrypted   bool        `json:"encrypted"`
	FrontendKey string      `json:"frontend_key"`
}

func main() {
	// Read schema file
	schemaData, err := os.ReadFile("internal/config/settings_schema.json")
	if err != nil {
		fmt.Printf("Error reading schema: %v\n", err)
		os.Exit(1)
	}

	var schema SettingsSchema
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		fmt.Printf("Error parsing schema: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🔧 Generating code from schema with %d settings...\n\n", len(schema.Settings))

	// Generate all files
	if err := generateDefaultsJSON(&schema); err != nil {
		fmt.Printf("Error generating defaults.json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Generated config/defaults.json")

	if err := generateInternalDefaultsJSON(&schema); err != nil {
		fmt.Printf("Error generating internal defaults.json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Generated internal/config/defaults.json")

	if err := generateConfigGo(&schema); err != nil {
		fmt.Printf("Error generating config.go: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Generated internal/config/config.go")

	if err := generateSettingsKeysGo(&schema); err != nil {
		fmt.Printf("Error generating settings_keys.go: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Generated internal/config/settings_keys.go")

	if err := generateSettingsHandlersGo(&schema); err != nil {
		fmt.Printf("Error generating settings_handlers.go: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Generated internal/handlers/settings/settings_handlers.go")

	if err := generateFrontendTypes(&schema); err != nil {
		fmt.Printf("Error generating frontend types: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Generated frontend/src/types/settings.generated.ts")

	if err := generateFrontendComposable(&schema); err != nil {
		fmt.Printf("Error generating frontend composable: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Generated frontend/src/composables/core/useSettings.generated.ts")

	fmt.Println("\n✨ All files generated successfully!")
	fmt.Println("\n📝 Next steps:")
	fmt.Println("1. Review generated files")
	fmt.Println("2. Run 'go build' to verify backend code")
	fmt.Println("3. Run 'cd frontend && npm run build' to verify frontend code")
	fmt.Println("4. Update database/db.go to use config.SettingsKeys()")
	fmt.Println("5. Test the application")
}

func generateDefaultsJSON(schema *SettingsSchema) error {
	defaults := make(map[string]interface{})
	// Use backend snake_case keys for defaults.json
	for key, def := range schema.Settings {
		defaults[key] = def.Default
	}

	data, err := json.MarshalIndent(defaults, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("config/defaults.json", data, 0644)
}

func generateInternalDefaultsJSON(schema *SettingsSchema) error {
	defaults := make(map[string]interface{})
	// Use backend snake_case keys for defaults.json
	for key, def := range schema.Settings {
		defaults[key] = def.Default
	}

	data, err := json.MarshalIndent(defaults, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("internal/config/defaults.json", data, 0644)
}

func generateConfigGo(schema *SettingsSchema) error {
	// Build struct fields and switch cases
	var structFields []string
	var switchCases []string

	for key, def := range schema.Settings {
		// Convert key to Go field name
		goKey := toGoFieldName(key)
		goType := toGoType(def.Type)

		// Struct field with JSON tag
		structFields = append(structFields, fmt.Sprintf("\t%s %s `json:\"%s\"`", goKey, goType, key))

		// Switch case for GetString
		caseStmt := fmt.Sprintf("\tcase \"%s\":", key)
		var returnValue string
		switch def.Type {
		case "int":
			returnValue = fmt.Sprintf("strconv.Itoa(defaults.%s)", goKey)
		case "bool":
			returnValue = fmt.Sprintf("strconv.FormatBool(defaults.%s)", goKey)
		case "string":
			returnValue = fmt.Sprintf("defaults.%s", goKey)
		}
		switchCases = append(switchCases, caseStmt, "\t\treturn "+returnValue)
	}

	tmpl := `// Copyright 2026 Ch3nyang & MrRSS Team. All rights reserved.
//
// Package config provides centralized default values for settings.
// The defaults are loaded from config/defaults.json which is shared between
// frontend and backend to ensure consistency.
// CODE GENERATED - DO NOT EDIT MANUALLY
// To add new settings, edit internal/config/settings_schema.json and run: go run tools/settings-generator/main.go
package config

import (
	_ "embed"
	"encoding/json"
	"strconv"
)

//go:embed defaults.json
var defaultsJSON []byte

// Defaults holds all default settings values
type Defaults struct {
%s
}

var defaults Defaults

func init() {
	if err := json.Unmarshal(defaultsJSON, &defaults); err != nil {
		panic("failed to parse defaults.json: " + err.Error())
	}
}

// Get returns the loaded defaults
func Get() Defaults {
	return defaults
}

// GetString returns a setting default as a string
func GetString(key string) string {
	switch key {
%s
	default:
		return ""
	}
}
`

	content := fmt.Sprintf(tmpl,
		strings.Join(structFields, "\n"),
		strings.Join(switchCases, "\n"))

	return os.WriteFile("internal/config/config.go", []byte(content), 0644)
}

func generateSettingsKeysGo(schema *SettingsSchema) error {
	var keys []string
	for key := range schema.Settings {
		keys = append(keys, fmt.Sprintf("\"%s\"", key))
	}

	tmpl := `// Copyright 2026 Ch3nyang & MrRSS Team. All rights reserved.
//
// Package config provides settings keys for database initialization
// CODE GENERATED - DO NOT EDIT MANUALLY
// To add new settings, edit internal/config/settings_schema.json and run: go run tools/settings-generator/main.go
package config

// SettingsKeys returns all valid setting keys
func SettingsKeys() []string {
	return []string{%s}
}
`

	content := fmt.Sprintf(tmpl, strings.Join(keys, ", "))
	return os.WriteFile("internal/config/settings_keys.go", []byte(content), 0644)
}

func generateSettingsHandlersGo(schema *SettingsSchema) error {
	// Generate GET variables
	var getVars []string
	var jsonFields []string
	for key, def := range schema.Settings {
		varName := toGoVarName(key)
		if def.Encrypted {
			getVars = append(getVars, fmt.Sprintf("\t%s, _ := h.DB.GetEncryptedSetting(\"%s\")", varName, key))
		} else {
			getVars = append(getVars, fmt.Sprintf("\t%s, _ := h.DB.GetSetting(\"%s\")", varName, key))
		}
		jsonFields = append(jsonFields, fmt.Sprintf("\t\t\"%s\": %s,", key, varName))
	}

	// Generate POST struct fields and save logic
	var structFields []string
	var saveStatements []string
	for key, def := range schema.Settings {
		goKey := toGoFieldName(key)
		structFields = append(structFields, fmt.Sprintf("\t\t%s string `json:\"%s\"`", goKey, key))

		if def.Encrypted {
			saveStatements = append(saveStatements, fmt.Sprintf(`		if err := h.DB.SetEncryptedSetting("%s", req.%s); err != nil {
			log.Printf("Failed to save %s: %%v", err)
			http.Error(w, "Failed to save %s", http.StatusInternalServerError)
			return
		}`, key, goKey, key, key))
		} else {
			saveStatements = append(saveStatements, fmt.Sprintf(`		if req.%s != "" {
			h.DB.SetSetting("%s", req.%s)
		}`, goKey, key, goKey))
		}
	}

	tmpl := `package settings

import (
	"encoding/json"
	"log"
	"net/http"

	"MrRSS/internal/handlers/core"
)

// HandleSettings handles GET and POST requests for application settings.
// CODE GENERATED - DO NOT EDIT MANUALLY
// To add new settings, edit internal/config/settings_schema.json and run: go run tools/settings-generator/main.go
func HandleSettings(h *core.Handler, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
%s
		json.NewEncoder(w).Encode(map[string]string{
%s
		})
	case http.MethodPost:
		var req struct {
%s
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
%s
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
`

	content := fmt.Sprintf(tmpl,
		strings.Join(getVars, "\n"),
		strings.Join(jsonFields, "\n"),
		strings.Join(structFields, "\n"),
		strings.Join(saveStatements, "\n\n"))

	return os.WriteFile("internal/handlers/settings/settings_handlers.go", []byte(content), 0644)
}

func generateFrontendTypes(schema *SettingsSchema) error {
	var fields []string
	for _, def := range schema.Settings {
		feKey := def.FrontendKey
		tsType := toTSType(def.Type)
		fields = append(fields, fmt.Sprintf("\t%s: %s;", feKey, tsType))
	}

	tmpl := `// Copyright 2026 Ch3nyang & MrRSS Team. All rights reserved.
//
// Auto-generated settings types
// CODE GENERATED - DO NOT EDIT MANUALLY
// To add new settings, edit internal/config/settings_schema.json and run: go run tools/settings-generator/main.go

export interface SettingsData {
%s
}
`

	content := fmt.Sprintf(tmpl, strings.Join(fields, "\n"))
	return os.WriteFile("frontend/src/types/settings.generated.ts", []byte(content), 0644)
}

func generateFrontendComposable(schema *SettingsSchema) error {
	var fetchFields []string
	var autoSaveFields []string
	var eventListeners []string

	for _, def := range schema.Settings {
		feKey := def.FrontendKey
		defaultVal := fmt.Sprintf("%#v", def.Default)
		backendKey := toBackendKey(feKey)

		// Fetch field conversion
		switch def.Type {
		case "bool":
			fetchFields = append(fetchFields, fmt.Sprintf("\t\t%s: data.%s === 'true',", feKey, backendKey))
		case "int":
			fetchFields = append(fetchFields, fmt.Sprintf("\t\t%s: parseInt(data.%s) || %s,", feKey, backendKey, defaultVal))
		default:
			fetchFields = append(fetchFields, fmt.Sprintf("\t\t%s: data.%s || %s,", feKey, backendKey, defaultVal))
		}

		// Auto-save field
		switch def.Type {
		case "bool":
			autoSaveFields = append(autoSaveFields, fmt.Sprintf("\t\t%s: (settingsRef.value.%s ?? settingsDefaults.%s).toString(),", backendKey, feKey, feKey))
		default:
			autoSaveFields = append(autoSaveFields, fmt.Sprintf("\t\t%s: settingsRef.value.%s ?? settingsDefaults.%s,", backendKey, feKey, feKey))
		}

		// Event listener (for non-internal settings)
		if def.Category != "internal" {
			eventName := toKebabCase(feKey) + "-changed"
			eventListeners = append(eventListeners, fmt.Sprintf("\t\twindow.dispatchEvent(new CustomEvent('%s', { detail: { value: settingsRef.value.%s } }))", eventName, feKey))
		}
	}

	tmpl := `// Copyright 2026 Ch3nyang & MrRSS Team. All rights reserved.
//
// Auto-generated settings composable functions
// CODE GENERATED - DO NOT EDIT MANUALLY
// To add new settings, edit internal/config/settings_schema.json and run: go run tools/settings-generator/main.go
import { type Ref } from 'vue'
import type { SettingsData } from '@/types/settings.generated'

// Generated settings defaults
export const settingsDefaults = {
%s
} as const

// Generated fetchSettings function
export function useGeneratedFetchSettings(settingsRef: Ref<SettingsData>) {
  return async () => {
    const response = await fetch('/api/settings')
    if (!response.ok) {
      throw new Error('Failed to fetch settings')
    }
    const data = await response.json()
    settingsRef.value = {
%s
    }
  }
}

// Generated auto-save payload builder
export function buildSavePayload(settingsRef: Ref<SettingsData>) {
  return {
%s
  }
}

// Generated event dispatchers
export function dispatchSettingChangeEvents(settingsRef: Ref<SettingsData>) {
%s
}
`

	// Build defaults object
	var defaultLines []string
	for _, def := range schema.Settings {
		feKey := def.FrontendKey
		defaultVal := fmt.Sprintf("%#v", def.Default)
		defaultLines = append(defaultLines, fmt.Sprintf("\t%s: %s,", feKey, defaultVal))
	}

	content := fmt.Sprintf(tmpl,
		strings.Join(defaultLines, "\n"),
		strings.Join(fetchFields, "\n"),
		strings.Join(autoSaveFields, "\n"),
		strings.Join(eventListeners, "\n"))

	return os.WriteFile("frontend/src/composables/core/useSettings.generated.ts", []byte(content), 0644)
}

// Helper functions
func toGoFieldName(key string) string {
	parts := strings.Split(key, "_")
	for i := 0; i < len(parts); i++ {
		// Capitalize first letter
		if len(parts[i]) > 0 {
			// For freshrss at start, make it FreshRSS
			if i == 0 && parts[i] == "freshrss" {
				parts[i] = "FreshRSS"
			} else if parts[i] == "ai" && i == 0 {
				// ai_ prefix at start should be AI
				parts[i] = "AI"
			} else if parts[i] == "ai" || parts[i] == "api" || parts[i] == "rss" {
				// Keep AI, API, RSS etc uppercase
				parts[i] = strings.ToUpper(parts[i])
			} else {
				// Capitalize first letter
				parts[i] = strings.ToUpper(string(parts[i][0])) + parts[i][1:]
			}
		}
	}
	return strings.Join(parts, "")
}

func toGoVarName(key string) string {
	// Handle freshrss specially
	if strings.HasPrefix(key, "freshrss") {
		parts := strings.Split(key, "_")
		for i := 1; i < len(parts); i++ {
			if len(parts[i]) > 0 {
				parts[i] = strings.ToUpper(string(parts[i][0])) + parts[i][1:]
			}
		}
		return strings.Join(parts, "")
	}

	// Handle multi-word keys
	parts := strings.Split(key, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(string(parts[i][0])) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

func toGoType(typ string) string {
	switch typ {
	case "int":
		return "int"
	case "bool":
		return "bool"
	case "string":
		return "string"
	default:
		return "string"
	}
}

func toTSType(typ string) string {
	switch typ {
	case "int":
		return "number"
	case "bool":
		return "boolean"
	case "string":
		return "string"
	default:
		return "string"
	}
}

func toKebabCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '-')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func toBackendKey(feKey string) string {
	// Convert camelCase frontend key to snake_case backend key
	var result []rune
	for i, r := range feKey {
		if r >= 'A' && r <= 'Z' {
			result = append(result, '_')
			result = append(result, r+32) // to lowercase
		} else {
			result = append(result, r)
		}
		if i == 0 && r >= 'A' && r <= 'Z' {
			// Keep first letter uppercase for FreshRSS etc
			result = result[:0]
			result = append(result, r)
		}
	}
	// Handle special cases
	resultStr := string(result)
	resultStr = strings.ReplaceAll(resultStr, "fresh_r_s_s", "freshrss")
	resultStr = strings.ReplaceAll(resultStr, "a_i", "ai")
	return resultStr
}
