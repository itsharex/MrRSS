# Settings Management System

## Table of Contents

- [Quick Start](#quick-start)
- [Overview](#overview)
- [How to Add a New Setting](#how-to-add-a-new-setting)
- [Complete Example](#complete-example)
- [Reference](#reference)

---

## Quick Start

**Adding a new setting is now simple - just 3 steps:**

### 1. Add to Schema

Edit `internal/config/settings_schema.json`:

```json
"your_setting_key": {
  "type": "bool",              // or "int", "string"
  "default": false,
  "category": "general",
  "encrypted": false,
  "frontend_key": "yourSettingKey"
}
```

### 2. Generate Code

```bash
go run tools/settings-generator/main.go
```

### 3. Add UI (Optional)

Add to your settings component:

```vue
<SettingItem :title="t('yourSettingKey')">
  <Toggle
    :model-value="settings.yourSettingKey"
    @update:model-value="updateSetting('yourSettingKey', $event)"
  />
</SettingItem>
```

That's it! See [Complete Example](#complete-example) for a detailed walkthrough.

---

## Overview

The settings system has been optimized to use **schema-driven code generation**. Instead of manually editing 11+ files, you now only need to edit **1 file** and run the code generator.

### Before vs After

**Old Way (Deprecated):**

- Edit 11 files manually
- ~50-100 lines of repetitive code
- High chance of copy-paste errors
- 30-45 minutes of work

**New Way (Current):**

- Edit 1 file (5 lines)
- Run 1 command
- Add UI and translations (optional)
- 10-15 minutes of work

**Result:** ~90% reduction in development time, near-zero error risk

### What Gets Generated

After running the generator, these files are **automatically created/updated**:

- ✅ `internal/config/config.go` - Go struct and `GetString()` function
- ✅ `internal/config/settings_keys.go` - Settings keys array for DB init
- ✅ `internal/handlers/settings/settings_handlers.go` - GET/POST API handlers
- ✅ `frontend/src/types/settings.generated.ts` - TypeScript interface
- ✅ `frontend/src/composables/core/useSettings.generated.ts` - Helper functions
- ✅ `config/defaults.json` - Frontend defaults (snake_case)
- ✅ `internal/config/defaults.json` - Backend defaults (snake_case)

---

## How to Add a New Setting

### Step 1: Define in Schema

Edit `internal/config/settings_schema.json`, add to the `settings` object:

```json
"your_new_setting": {
  "type": "bool",              // Type: "bool", "int", or "string"
  "default": false,            // Default value
  "category": "general",       // Category (see reference below)
  "encrypted": false,          // Set to true for sensitive data
  "frontend_key": "yourNewSetting"  // camelCase version for frontend
}
```

**Schema Properties:**

| Property | Type | Description |
| -------- | ---- | ----------- |
| `type` | string | Required: `"bool"`, `"int"`, or `"string"` |
| `default` | mixed | Required: Default value (must match type) |
| `category` | string | Required: See [Categories](#categories) below |
| `encrypted` | boolean | Required: `true` for sensitive data (API keys, passwords) |
| `frontend_key` | string | Required: camelCase version for frontend |

### Step 2: Run the Code Generator

```bash
go run tools/settings-generator/main.go
```

**Output:**

```plaintext
🔧 Generating code from schema with 66 settings...

✓ Generated config/defaults.json
✓ Generated internal/config/defaults.json
✓ Generated internal/config/config.go
✓ Generated internal/config/settings_keys.go
✓ Generated internal/handlers/settings/settings_handlers.go
✓ Generated frontend/src/types/settings.generated.ts
✓ Generated frontend/src/composables/core/useSettings.generated.ts

✨ All files generated successfully!
```

This automatically generates all the boilerplate code for both backend and frontend.

### Step 3: Add Translations (Recommended)

#### English (`frontend/src/i18n/locales/en.ts`)

Find the appropriate section and add:

```typescript
yourNewSetting: 'Your New Setting',
yourNewSettingDesc: 'Description of what this setting does',
```

#### Chinese (`frontend/src/i18n/locales/zh.ts`)

```typescript
yourNewSetting: '您的新设置',
yourNewSettingDesc: '此设置功能的描述',
```

### Step 4: Add UI (Optional)

Add the setting UI to the appropriate settings component.

**Example** - `frontend/src/components/modals/settings/general/GeneralSettings.vue`:

```vue
<SettingItem
  :title="t('yourNewSetting')"
  :description="t('yourNewSettingDesc')"
>
  <Toggle
    :model-value="settings.yourNewSetting"
    @update:model-value="updateSetting('yourNewSetting', $event)"
  />
</SettingItem>
```

**UI Component Examples:**

```vue
<!-- Boolean/Toggle -->
<Toggle
  :model-value="settings.yourSetting"
  @update:model-value="updateSetting('yourSetting', $event)"
/>

<!-- String/Input -->
<Input
  v-model="settings.yourSetting"
  @change="updateSetting('yourSetting', $event)"
/>

<!-- Integer/Number -->
<Input
  v-model.number="settings.yourSetting"
  type="number"
  @change="updateSetting('yourSetting', $event)"
/>

<!-- Select/Enum -->
<Select
  v-model="settings.yourSetting"
  :options="[{value: 'option1', label: 'Option 1'}, ...]"
  @change="updateSetting('yourSetting', $event)"
/>
```

### Step 5: Implement Feature Logic (Optional)

If the setting affects app behavior, implement the logic.

#### Option A: Listen to Settings Event

```vue
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useSettings } from '@/composables/core/useSettings'

const { settings } = useSettings()
const featureEnabled = ref(false)

onMounted(() => {
  // Apply the setting
  featureEnabled.value = settings.value.yourSetting
})

// Listen for changes
window.addEventListener('your-setting-changed', (event: any) => {
  featureEnabled.value = event.detail.value
})
</script>
```

#### Option B: Use Composable

Create `frontend/src/composables/core/useYourFeature.ts`:

```typescript
import { computed } from 'vue'
import { useSettings } from './useSettings'

export function useYourFeature() {
  const { settings } = useSettings()

  const featureEnabled = computed(() => settings.value.yourSetting)

  return {
    featureEnabled
  }
}
```

### Step 6: Test

```bash
# Backend
go build

# Frontend
cd frontend
npm run build

# Or run full dev mode
cd ..
wails3 dev
```

---

## Complete Example

Let's walk through adding a complete setting from start to finish.

### Goal: Add "Auto-Collapse Sidebar" Setting

We want to add a setting that automatically collapses the sidebar on startup.

### Step 1: Define in Schema

Edit `internal/config/settings_schema.json`:

```json
"auto_collapse_sidebar": {
  "type": "bool",
  "default": false,
  "category": "general",
  "encrypted": false,
  "frontend_key": "autoCollapseSidebar"
}
```

**Why these values?**

- `type: "bool"` - It's a toggle/checkbox setting
- `default: false` - Most users want sidebar expanded by default
- `category: "general"` - It's a general UI preference
- `encrypted: false` - Not sensitive data
- `frontend_key: "autoCollapseSidebar"` - camelCase for frontend

### Step 2: Generate Code

```bash
go run tools/settings-generator/main.go
```

**What was generated?**

1. **`internal/config/config.go`**
   - Added `AutoCollapseSidebar bool` field to Defaults struct
   - Added switch case for `GetString("auto_collapse_sidebar")`

2. **`internal/config/settings_keys.go`**
   - Added `"auto_collapse_sidebar"` to keys array

3. **`internal/handlers/settings/settings_handlers.go`**
   - Added GET: `autoCollapseSidebar, _ := h.DB.GetSetting("auto_collapse_sidebar")`
   - Added JSON response: `"auto_collapse_sidebar": autoCollapseSidebar`
   - Added POST field: `AutoCollapseSidebar string \`json:"auto_collapse_sidebar"\``
   - Added save logic: `if req.AutoCollapseSidebar != "" { h.DB.SetSetting(...) }`

4. **`frontend/src/types/settings.generated.ts`**
   - Added: `autoCollapseSidebar: boolean;`

5. **`frontend/src/composables/core/useSettings.generated.ts`**
   - Added: `autoCollapseSidebar: false,` to defaults
   - Added fetch: `autoCollapseSidebar: data.auto_collapse_sidebar === 'true',`
   - Added save: `auto_collapse_sidebar: (settingsRef.value.autoCollapseSidebar ?? settingsDefaults.autoCollapseSidebar).toString(),`
   - Added event: `window.dispatchEvent(new CustomEvent('auto-collapse-sidebar-changed', ...))`

6. **`config/defaults.json` & `internal/config/defaults.json`**
   - Added: `"auto_collapse_sidebar": false`

### Step 3: Add Translations

**English** (`frontend/src/i18n/locales/en.ts`):

```typescript
autoCollapseSidebar: 'Auto Collapse Sidebar',
autoCollapseSidebarDesc: 'Automatically collapse the sidebar when the app starts',
```

**Chinese** (`frontend/src/i18n/locales/zh.ts`):

```typescript
autoCollapseSidebar: '自动折叠侧边栏',
autoCollapseSidebarDesc: '应用启动时自动折叠侧边栏',
```

### Step 4: Add UI

Add to `frontend/src/components/modals/settings/general/GeneralSettings.vue`:

```vue
<SettingItem
  :title="t('autoCollapseSidebar')"
  :description="t('autoCollapseSidebarDesc')"
>
  <Toggle
    :model-value="settings.autoCollapseSidebar"
    @update:model-value="updateSetting('autoCollapseSidebar', $event)"
  />
</SettingItem>
```

Place it near related settings (like theme, startup on boot).

### Step 5: Implement Feature Logic

In your sidebar component:

```vue
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useSettings } from '@/composables/core/useSettings'

const { settings } = useSettings()
const isCollapsed = ref(false)

onMounted(() => {
  // Apply the setting
  isCollapsed.value = settings.value.autoCollapseSidebar
})

// Listen for changes
window.addEventListener('auto-collapse-sidebar-changed', (event: any) => {
  isCollapsed.value = event.detail.value
})
</script>

<template>
  <aside :class="{ collapsed: isCollapsed }">
    <!-- Sidebar content -->
  </aside>
</template>

<style scoped>
aside.collapsed {
  width: 60px;
}
</style>
```

### Step 6: Test

#### Manual Testing

1. Open Settings → General
2. Find "Auto Collapse Sidebar" setting
3. Toggle it on
4. Close and reopen the app
5. ✅ Verify sidebar is collapsed on startup
6. Toggle it off
7. Close and reopen the app
8. ✅ Verify sidebar is expanded on startup

#### Check Database

```sql
SELECT * FROM settings WHERE key = 'auto_collapse_sidebar';
```

Should show:

```plaintext
key                     | value
------------------------+-------
auto_collapse_sidebar   | true
```

#### Verify API

**GET** `/api/settings`:

```bash
curl http://localhost:5343/api/settings
```

Should include:

```json
{
  "auto_collapse_sidebar": "true"
}
```

**POST** `/api/settings`:

```bash
curl -X POST http://localhost:5343/api/settings \
  -H "Content-Type: application/json" \
  -d '{"auto_collapse_sidebar": "false"}'
```

Should return `200 OK`.

### Complete Checklist

- [x] Schema added to `settings_schema.json`
- [x] Code generator ran successfully
- [x] Backend compiles without errors
- [x] Frontend compiles without errors
- [x] English translations added
- [x] Chinese translations added
- [x] UI component added (Toggle in GeneralSettings)
- [x] Feature logic implemented (sidebar collapse)
- [x] Setting appears in settings modal
- [x] Setting saves to database
- [x] Setting loads on startup
- [x] API GET returns correct value
- [x] API POST saves value correctly

---

## Reference

### Type Mapping

| Schema Type | Go Type | TypeScript Type | Example |
| ----------- | ------- | --------------- | ------- |
| `"bool"` | `bool` | `boolean` | `true`, `false` |
| `"int"` | `int` | `number` | `30`, `500` |
| `"string"` | `string` | `string` | `"en"`, `"openai"` |

### Categories

Use the appropriate category for your setting:

| Category | Description | Example Settings |
| -------- | ----------- | ---------------- |
| `general` | General app settings | theme, language, shortcuts |
| `reading` | Reading/viewing preferences | view mode, hover mark as read, show hidden |
| `translation` | Translation settings | provider, target language, API keys |
| `ai` | AI-related settings | API key, model, prompts, usage limit |
| `summary` | Article summary settings | summary length, trigger mode |
| `storage` | Cache and storage settings | cache size, cleanup, max age |
| `network` | Network and proxy settings | proxy, bandwidth, concurrent refreshes |
| `integrations` | Third-party integrations | Obsidian, FreshRSS |
| `internal` | Internal app state (no UI) | window position, last update |

### Encrypted Settings

For sensitive data (API keys, passwords), set `"encrypted": true`:

```json
"my_api_key": {
  "type": "string",
  "default": "",
  "category": "integrations",
  "encrypted": true,  // ← Important!
  "frontend_key": "myAPIKey"
}
```

Encrypted settings are automatically:

- Stored encrypted in the database
- Fetched using `GetEncryptedSetting()` instead of `GetSetting()`
- Saved using `SetEncryptedSetting()` instead of `SetSetting()`

### Frontend Key Naming

The `frontend_key` should be in **camelCase** and map to the backend key:

| Backend Key | Frontend Key |
| ----------- | ----------- |
| `update_interval` | `updateInterval` |
| `startup_on_boot` | `startupOnBoot` |
| `deepl_api_key` | `deeplAPIKey` |
| `ai_endpoint` | `aiEndpoint` |
| `freshrss_enabled` | `freshRSSSyncEnabled` |

### Quick Examples

**Boolean Setting:**

```json
"enable_feature": {
  "type": "bool",
  "default": true,
  "category": "general",
  "encrypted": false,
  "frontend_key": "enableFeature"
}
```

**Integer Setting:**

```json
"max_items": {
  "type": "int",
  "default": 100,
  "category": "storage",
  "encrypted": false,
  "frontend_key": "maxItems"
}
```

**String Setting:**

```json
"api_endpoint": {
  "type": "string",
  "default": "https://api.example.com",
  "category": "integrations",
  "encrypted": false,
  "frontend_key": "apiEndpoint"
}
```

**Encrypted Setting:**

```json
"api_secret": {
  "type": "string",
  "default": "",
  "category": "integrations",
  "encrypted": true,    // ← Encrypts in DB
  "frontend_key": "apiSecret"
}
```

### Event Listeners

For non-internal settings, change events are automatically dispatched. Listen to them like this:

```typescript
window.addEventListener('your-setting-key-changed', (event) => {
  const { value } = event.detail
  console.log('Setting changed to:', value)
})
```

Event name format: `{frontend_key in kebab-case}-changed`

Examples:

- `autoCollapseSidebar` → `auto-collapse-sidebar-changed`
- `aiAPIKey` → `ai-api-key-changed`
- `freshRSSSyncEnabled` → `fresh-rss-sync-enabled-changed`

### Common Mistakes

❌ **Wrong:**

```json
"type": "boolean",     // Should be "bool"
"category": "General", // Should be lowercase
"frontend_key": "my_setting" // Should be camelCase
```

✅ **Correct:**

```json
"type": "bool",
"category": "general",
"frontend_key": "mySetting"
```

### Troubleshooting

#### Build Errors

**Problem:** `go build` fails after adding a setting

**Solution:**

1. Check that your `settings_schema.json` has valid JSON syntax (no missing commas)
2. Verify `type` is one of: `"bool"`, `"int"`, `"string"`
3. Verify `category` is a valid category
4. Run generator again: `go run tools/settings-generator/main.go`

#### Frontend Errors

**Problem:** `Property 'mySetting' does not exist`

**Solution:**

1. Make sure you ran the generator
2. Check `frontend/src/types/settings.generated.ts` exists and has your setting
3. Try `npm run build` in frontend directory
4. Restart TypeScript server in VSCode

#### Setting Not Appearing in UI

**Problem:** Toggle doesn't show in settings modal

**Solution:**

1. Check that you added the `<SettingItem>` component
2. Verify the translation keys match
3. Check browser console for errors
4. Try hard refresh (Ctrl+Shift+R)

#### Setting Not Saving

**Problem:** Toggle changes but resets on restart

**Solution:**

1. Open browser DevTools → Network tab
2. Check if POST to `/api/settings` is sent
3. Check response status (should be 200 OK)
4. Check database directly via SQLite browser
5. Verify `frontend_key` matches in schema

---

## Migration from Old System

If you have existing manually-written settings code:

1. ✅ Ensure all settings are defined in `internal/config/settings_schema.json`
2. ✅ Run the generator: `go run tools/settings-generator/main.go`
3. ✅ Review and commit the generated files
4. ✅ Delete any manual setting-related code that's now replaced

The generated code is compatible with the existing database and API.

---

## Best Practices

1. **Use descriptive names** - `enable_auto_sync` not `eas`
2. **Choose appropriate types** - Use `bool` for toggles, `int` for numbers
3. **Set sensible defaults** - What should the setting be for new users?
4. **Add translations** - Always add both English and Chinese
5. **Use categories** - This helps organize the settings UI
6. **Encrypt sensitive data** - API keys, passwords, tokens
7. **Test after adding** - Run the app and verify the setting works
8. **Document complex settings** - Add comments if behavior is non-obvious

---

## Summary

**Old workflow:** Edit 11 files, ~100 lines of code, high chance of errors

**New workflow:** Edit 1 file (5 lines), run 1 command, done!

This optimization:

- ✅ Reduces development time by ~90%
- ✅ Eliminates copy-paste errors
- ✅ Ensures consistency between frontend and backend
- ✅ Maintains type safety automatically
- ✅ Makes adding new settings trivial

Happy coding! 🚀
