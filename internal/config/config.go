// Copyright 2026 Ch3nyang & MrRSS Team. All rights reserved.
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
	FullTextFetchEnabled     bool   `json:"full_text_fetch_enabled"`
	Theme                    string `json:"theme"`
	AITranslationPrompt      string `json:"ai_translation_prompt"`
	ObsidianVault            string `json:"obsidian_vault"`
	Language                 string `json:"language"`
	AIAPIKey                 string `json:"ai_api_key"`
	MediaCacheEnabled        bool   `json:"media_cache_enabled"`
	ProxyUsername            string `json:"proxy_username"`
	MaxConcurrentRefreshes   string `json:"max_concurrent_refreshes"`
	AISummaryPrompt          string `json:"ai_summary_prompt"`
	ProxyPassword            string `json:"proxy_password"`
	LastArticleUpdate        string `json:"last_article_update"`
	LastNetworkTest          string `json:"last_network_test"`
	ImageGalleryEnabled      bool   `json:"image_gallery_enabled"`
	UpdateInterval           int    `json:"update_interval"`
	AIUsageLimit             string `json:"ai_usage_limit"`
	AIChatEnabled            bool   `json:"ai_chat_enabled"`
	SummaryTriggerMode       string `json:"summary_trigger_mode"`
	WindowHeight             string `json:"window_height"`
	WindowMaximized          string `json:"window_maximized"`
	FreshRSSEnabled          bool   `json:"freshrss_enabled"`
	RefreshMode              string `json:"refresh_mode"`
	TranslationEnabled       bool   `json:"translation_enabled"`
	TargetLanguage           string `json:"target_language"`
	ProxyEnabled             bool   `json:"proxy_enabled"`
	AutoShowAllContent       bool   `json:"auto_show_all_content"`
	AIUsageTokens            string `json:"ai_usage_tokens"`
	StartupOnBoot            bool   `json:"startup_on_boot"`
	CloseToTray              bool   `json:"close_to_tray"`
	DeeplAPIKey              string `json:"deepl_api_key"`
	DeeplEndpoint            string `json:"deepl_endpoint"`
	BaiduAppId               string `json:"baidu_app_id"`
	BaiduSecretKey           string `json:"baidu_secret_key"`
	MaxCacheSizeMb           int    `json:"max_cache_size_mb"`
	DefaultViewMode          string `json:"default_view_mode"`
	ShowHiddenArticles       bool   `json:"show_hidden_articles"`
	SummaryEnabled           bool   `json:"summary_enabled"`
	MediaCacheMaxSizeMb      int    `json:"media_cache_max_size_mb"`
	ProxyType                string `json:"proxy_type"`
	Rules                    string `json:"rules"`
	WindowWidth              string `json:"window_width"`
	NetworkSpeed             string `json:"network_speed"`
	Shortcuts                string `json:"shortcuts"`
	NetworkBandwidthMbps     string `json:"network_bandwidth_mbps"`
	FreshRSSUsername         string `json:"freshrss_username"`
	TranslationProvider      string `json:"translation_provider"`
	ProxyHost                string `json:"proxy_host"`
	WindowX                  string `json:"window_x"`
	FreshRSSServerUrl        string `json:"freshrss_server_url"`
	MediaCacheMaxAgeDays     int    `json:"media_cache_max_age_days"`
	GoogleTranslateEndpoint  string `json:"google_translate_endpoint"`
	ShowArticlePreviewImages bool   `json:"show_article_preview_images"`
	NetworkLatencyMs         string `json:"network_latency_ms"`
	SummaryLength            string `json:"summary_length"`
	MaxArticleAgeDays        int    `json:"max_article_age_days"`
	ProxyPort                string `json:"proxy_port"`
	HoverMarkAsRead          bool   `json:"hover_mark_as_read"`
	AIEndpoint               string `json:"ai_endpoint"`
	ObsidianVaultPath        string `json:"obsidian_vault_path"`
	AIModel                  string `json:"ai_model"`
	WindowY                  string `json:"window_y"`
	AICustomHeaders          string `json:"ai_custom_headers"`
	AutoCleanupEnabled       bool   `json:"auto_cleanup_enabled"`
	ObsidianEnabled          bool   `json:"obsidian_enabled"`
	SummaryProvider          string `json:"summary_provider"`
	FreshRSSAPIPassword      string `json:"freshrss_api_password"`
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
	case "full_text_fetch_enabled":
		return strconv.FormatBool(defaults.FullTextFetchEnabled)
	case "theme":
		return defaults.Theme
	case "ai_translation_prompt":
		return defaults.AITranslationPrompt
	case "obsidian_vault":
		return defaults.ObsidianVault
	case "language":
		return defaults.Language
	case "ai_api_key":
		return defaults.AIAPIKey
	case "media_cache_enabled":
		return strconv.FormatBool(defaults.MediaCacheEnabled)
	case "proxy_username":
		return defaults.ProxyUsername
	case "max_concurrent_refreshes":
		return defaults.MaxConcurrentRefreshes
	case "ai_summary_prompt":
		return defaults.AISummaryPrompt
	case "proxy_password":
		return defaults.ProxyPassword
	case "last_article_update":
		return defaults.LastArticleUpdate
	case "last_network_test":
		return defaults.LastNetworkTest
	case "image_gallery_enabled":
		return strconv.FormatBool(defaults.ImageGalleryEnabled)
	case "update_interval":
		return strconv.Itoa(defaults.UpdateInterval)
	case "ai_usage_limit":
		return defaults.AIUsageLimit
	case "ai_chat_enabled":
		return strconv.FormatBool(defaults.AIChatEnabled)
	case "summary_trigger_mode":
		return defaults.SummaryTriggerMode
	case "window_height":
		return defaults.WindowHeight
	case "window_maximized":
		return defaults.WindowMaximized
	case "freshrss_enabled":
		return strconv.FormatBool(defaults.FreshRSSEnabled)
	case "refresh_mode":
		return defaults.RefreshMode
	case "translation_enabled":
		return strconv.FormatBool(defaults.TranslationEnabled)
	case "target_language":
		return defaults.TargetLanguage
	case "proxy_enabled":
		return strconv.FormatBool(defaults.ProxyEnabled)
	case "auto_show_all_content":
		return strconv.FormatBool(defaults.AutoShowAllContent)
	case "ai_usage_tokens":
		return defaults.AIUsageTokens
	case "startup_on_boot":
		return strconv.FormatBool(defaults.StartupOnBoot)
	case "close_to_tray":
		return strconv.FormatBool(defaults.CloseToTray)
	case "deepl_api_key":
		return defaults.DeeplAPIKey
	case "deepl_endpoint":
		return defaults.DeeplEndpoint
	case "baidu_app_id":
		return defaults.BaiduAppId
	case "baidu_secret_key":
		return defaults.BaiduSecretKey
	case "max_cache_size_mb":
		return strconv.Itoa(defaults.MaxCacheSizeMb)
	case "default_view_mode":
		return defaults.DefaultViewMode
	case "show_hidden_articles":
		return strconv.FormatBool(defaults.ShowHiddenArticles)
	case "summary_enabled":
		return strconv.FormatBool(defaults.SummaryEnabled)
	case "media_cache_max_size_mb":
		return strconv.Itoa(defaults.MediaCacheMaxSizeMb)
	case "proxy_type":
		return defaults.ProxyType
	case "rules":
		return defaults.Rules
	case "window_width":
		return defaults.WindowWidth
	case "network_speed":
		return defaults.NetworkSpeed
	case "shortcuts":
		return defaults.Shortcuts
	case "network_bandwidth_mbps":
		return defaults.NetworkBandwidthMbps
	case "freshrss_username":
		return defaults.FreshRSSUsername
	case "translation_provider":
		return defaults.TranslationProvider
	case "proxy_host":
		return defaults.ProxyHost
	case "window_x":
		return defaults.WindowX
	case "freshrss_server_url":
		return defaults.FreshRSSServerUrl
	case "media_cache_max_age_days":
		return strconv.Itoa(defaults.MediaCacheMaxAgeDays)
	case "google_translate_endpoint":
		return defaults.GoogleTranslateEndpoint
	case "show_article_preview_images":
		return strconv.FormatBool(defaults.ShowArticlePreviewImages)
	case "network_latency_ms":
		return defaults.NetworkLatencyMs
	case "summary_length":
		return defaults.SummaryLength
	case "max_article_age_days":
		return strconv.Itoa(defaults.MaxArticleAgeDays)
	case "proxy_port":
		return defaults.ProxyPort
	case "hover_mark_as_read":
		return strconv.FormatBool(defaults.HoverMarkAsRead)
	case "ai_endpoint":
		return defaults.AIEndpoint
	case "obsidian_vault_path":
		return defaults.ObsidianVaultPath
	case "ai_model":
		return defaults.AIModel
	case "window_y":
		return defaults.WindowY
	case "ai_custom_headers":
		return defaults.AICustomHeaders
	case "auto_cleanup_enabled":
		return strconv.FormatBool(defaults.AutoCleanupEnabled)
	case "obsidian_enabled":
		return strconv.FormatBool(defaults.ObsidianEnabled)
	case "summary_provider":
		return defaults.SummaryProvider
	case "freshrss_api_password":
		return defaults.FreshRSSAPIPassword
	default:
		return ""
	}
}
