package settings

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
		startupOnBoot, _ := h.DB.GetSetting("startup_on_boot")
		closeToTray, _ := h.DB.GetSetting("close_to_tray")
		deeplApiKey, _ := h.DB.GetEncryptedSetting("deepl_api_key")
		deeplEndpoint, _ := h.DB.GetSetting("deepl_endpoint")
		baiduAppId, _ := h.DB.GetSetting("baidu_app_id")
		baiduSecretKey, _ := h.DB.GetEncryptedSetting("baidu_secret_key")
		maxCacheSizeMb, _ := h.DB.GetSetting("max_cache_size_mb")
		defaultViewMode, _ := h.DB.GetSetting("default_view_mode")
		showHiddenArticles, _ := h.DB.GetSetting("show_hidden_articles")
		summaryEnabled, _ := h.DB.GetSetting("summary_enabled")
		mediaCacheMaxSizeMb, _ := h.DB.GetSetting("media_cache_max_size_mb")
		proxyType, _ := h.DB.GetSetting("proxy_type")
		rules, _ := h.DB.GetSetting("rules")
		windowWidth, _ := h.DB.GetSetting("window_width")
		networkSpeed, _ := h.DB.GetSetting("network_speed")
		shortcuts, _ := h.DB.GetSetting("shortcuts")
		networkBandwidthMbps, _ := h.DB.GetSetting("network_bandwidth_mbps")
		freshrssUsername, _ := h.DB.GetSetting("freshrss_username")
		translationProvider, _ := h.DB.GetSetting("translation_provider")
		proxyHost, _ := h.DB.GetSetting("proxy_host")
		windowX, _ := h.DB.GetSetting("window_x")
		freshrssServerUrl, _ := h.DB.GetSetting("freshrss_server_url")
		mediaCacheMaxAgeDays, _ := h.DB.GetSetting("media_cache_max_age_days")
		googleTranslateEndpoint, _ := h.DB.GetSetting("google_translate_endpoint")
		showArticlePreviewImages, _ := h.DB.GetSetting("show_article_preview_images")
		networkLatencyMs, _ := h.DB.GetSetting("network_latency_ms")
		summaryLength, _ := h.DB.GetSetting("summary_length")
		maxArticleAgeDays, _ := h.DB.GetSetting("max_article_age_days")
		proxyPort, _ := h.DB.GetSetting("proxy_port")
		hoverMarkAsRead, _ := h.DB.GetSetting("hover_mark_as_read")
		aiEndpoint, _ := h.DB.GetSetting("ai_endpoint")
		obsidianVaultPath, _ := h.DB.GetSetting("obsidian_vault_path")
		aiModel, _ := h.DB.GetSetting("ai_model")
		windowY, _ := h.DB.GetSetting("window_y")
		aiCustomHeaders, _ := h.DB.GetSetting("ai_custom_headers")
		autoCleanupEnabled, _ := h.DB.GetSetting("auto_cleanup_enabled")
		obsidianEnabled, _ := h.DB.GetSetting("obsidian_enabled")
		summaryProvider, _ := h.DB.GetSetting("summary_provider")
		freshrssApiPassword, _ := h.DB.GetEncryptedSetting("freshrss_api_password")
		fullTextFetchEnabled, _ := h.DB.GetSetting("full_text_fetch_enabled")
		theme, _ := h.DB.GetSetting("theme")
		aiTranslationPrompt, _ := h.DB.GetSetting("ai_translation_prompt")
		obsidianVault, _ := h.DB.GetSetting("obsidian_vault")
		language, _ := h.DB.GetSetting("language")
		aiApiKey, _ := h.DB.GetEncryptedSetting("ai_api_key")
		mediaCacheEnabled, _ := h.DB.GetSetting("media_cache_enabled")
		proxyUsername, _ := h.DB.GetEncryptedSetting("proxy_username")
		maxConcurrentRefreshes, _ := h.DB.GetSetting("max_concurrent_refreshes")
		aiSummaryPrompt, _ := h.DB.GetSetting("ai_summary_prompt")
		proxyPassword, _ := h.DB.GetEncryptedSetting("proxy_password")
		lastArticleUpdate, _ := h.DB.GetSetting("last_article_update")
		lastNetworkTest, _ := h.DB.GetSetting("last_network_test")
		imageGalleryEnabled, _ := h.DB.GetSetting("image_gallery_enabled")
		updateInterval, _ := h.DB.GetSetting("update_interval")
		aiUsageLimit, _ := h.DB.GetSetting("ai_usage_limit")
		aiChatEnabled, _ := h.DB.GetSetting("ai_chat_enabled")
		summaryTriggerMode, _ := h.DB.GetSetting("summary_trigger_mode")
		windowHeight, _ := h.DB.GetSetting("window_height")
		windowMaximized, _ := h.DB.GetSetting("window_maximized")
		freshrssEnabled, _ := h.DB.GetSetting("freshrss_enabled")
		refreshMode, _ := h.DB.GetSetting("refresh_mode")
		translationEnabled, _ := h.DB.GetSetting("translation_enabled")
		targetLanguage, _ := h.DB.GetSetting("target_language")
		proxyEnabled, _ := h.DB.GetSetting("proxy_enabled")
		autoShowAllContent, _ := h.DB.GetSetting("auto_show_all_content")
		aiUsageTokens, _ := h.DB.GetSetting("ai_usage_tokens")
		json.NewEncoder(w).Encode(map[string]string{
			"startup_on_boot":             startupOnBoot,
			"close_to_tray":               closeToTray,
			"deepl_api_key":               deeplApiKey,
			"deepl_endpoint":              deeplEndpoint,
			"baidu_app_id":                baiduAppId,
			"baidu_secret_key":            baiduSecretKey,
			"max_cache_size_mb":           maxCacheSizeMb,
			"default_view_mode":           defaultViewMode,
			"show_hidden_articles":        showHiddenArticles,
			"summary_enabled":             summaryEnabled,
			"media_cache_max_size_mb":     mediaCacheMaxSizeMb,
			"proxy_type":                  proxyType,
			"rules":                       rules,
			"window_width":                windowWidth,
			"network_speed":               networkSpeed,
			"shortcuts":                   shortcuts,
			"network_bandwidth_mbps":      networkBandwidthMbps,
			"freshrss_username":           freshrssUsername,
			"translation_provider":        translationProvider,
			"proxy_host":                  proxyHost,
			"window_x":                    windowX,
			"freshrss_server_url":         freshrssServerUrl,
			"media_cache_max_age_days":    mediaCacheMaxAgeDays,
			"google_translate_endpoint":   googleTranslateEndpoint,
			"show_article_preview_images": showArticlePreviewImages,
			"network_latency_ms":          networkLatencyMs,
			"summary_length":              summaryLength,
			"max_article_age_days":        maxArticleAgeDays,
			"proxy_port":                  proxyPort,
			"hover_mark_as_read":          hoverMarkAsRead,
			"ai_endpoint":                 aiEndpoint,
			"obsidian_vault_path":         obsidianVaultPath,
			"ai_model":                    aiModel,
			"window_y":                    windowY,
			"ai_custom_headers":           aiCustomHeaders,
			"auto_cleanup_enabled":        autoCleanupEnabled,
			"obsidian_enabled":            obsidianEnabled,
			"summary_provider":            summaryProvider,
			"freshrss_api_password":       freshrssApiPassword,
			"full_text_fetch_enabled":     fullTextFetchEnabled,
			"theme":                       theme,
			"ai_translation_prompt":       aiTranslationPrompt,
			"obsidian_vault":              obsidianVault,
			"language":                    language,
			"ai_api_key":                  aiApiKey,
			"media_cache_enabled":         mediaCacheEnabled,
			"proxy_username":              proxyUsername,
			"max_concurrent_refreshes":    maxConcurrentRefreshes,
			"ai_summary_prompt":           aiSummaryPrompt,
			"proxy_password":              proxyPassword,
			"last_article_update":         lastArticleUpdate,
			"last_network_test":           lastNetworkTest,
			"image_gallery_enabled":       imageGalleryEnabled,
			"update_interval":             updateInterval,
			"ai_usage_limit":              aiUsageLimit,
			"ai_chat_enabled":             aiChatEnabled,
			"summary_trigger_mode":        summaryTriggerMode,
			"window_height":               windowHeight,
			"window_maximized":            windowMaximized,
			"freshrss_enabled":            freshrssEnabled,
			"refresh_mode":                refreshMode,
			"translation_enabled":         translationEnabled,
			"target_language":             targetLanguage,
			"proxy_enabled":               proxyEnabled,
			"auto_show_all_content":       autoShowAllContent,
			"ai_usage_tokens":             aiUsageTokens,
		})
	case http.MethodPost:
		var req struct {
			MediaCacheMaxAgeDays     string `json:"media_cache_max_age_days"`
			GoogleTranslateEndpoint  string `json:"google_translate_endpoint"`
			ShowArticlePreviewImages string `json:"show_article_preview_images"`
			NetworkLatencyMs         string `json:"network_latency_ms"`
			SummaryLength            string `json:"summary_length"`
			MaxArticleAgeDays        string `json:"max_article_age_days"`
			ProxyPort                string `json:"proxy_port"`
			HoverMarkAsRead          string `json:"hover_mark_as_read"`
			AIEndpoint               string `json:"ai_endpoint"`
			ObsidianVaultPath        string `json:"obsidian_vault_path"`
			AIModel                  string `json:"ai_model"`
			WindowY                  string `json:"window_y"`
			AICustomHeaders          string `json:"ai_custom_headers"`
			AutoCleanupEnabled       string `json:"auto_cleanup_enabled"`
			ObsidianEnabled          string `json:"obsidian_enabled"`
			SummaryProvider          string `json:"summary_provider"`
			FreshRSSAPIPassword      string `json:"freshrss_api_password"`
			FullTextFetchEnabled     string `json:"full_text_fetch_enabled"`
			Theme                    string `json:"theme"`
			AITranslationPrompt      string `json:"ai_translation_prompt"`
			ObsidianVault            string `json:"obsidian_vault"`
			Language                 string `json:"language"`
			AIAPIKey                 string `json:"ai_api_key"`
			MediaCacheEnabled        string `json:"media_cache_enabled"`
			ProxyUsername            string `json:"proxy_username"`
			MaxConcurrentRefreshes   string `json:"max_concurrent_refreshes"`
			AISummaryPrompt          string `json:"ai_summary_prompt"`
			ProxyPassword            string `json:"proxy_password"`
			LastArticleUpdate        string `json:"last_article_update"`
			LastNetworkTest          string `json:"last_network_test"`
			ImageGalleryEnabled      string `json:"image_gallery_enabled"`
			UpdateInterval           string `json:"update_interval"`
			AIUsageLimit             string `json:"ai_usage_limit"`
			AIChatEnabled            string `json:"ai_chat_enabled"`
			SummaryTriggerMode       string `json:"summary_trigger_mode"`
			WindowHeight             string `json:"window_height"`
			WindowMaximized          string `json:"window_maximized"`
			FreshRSSEnabled          string `json:"freshrss_enabled"`
			RefreshMode              string `json:"refresh_mode"`
			TranslationEnabled       string `json:"translation_enabled"`
			TargetLanguage           string `json:"target_language"`
			ProxyEnabled             string `json:"proxy_enabled"`
			AutoShowAllContent       string `json:"auto_show_all_content"`
			AIUsageTokens            string `json:"ai_usage_tokens"`
			StartupOnBoot            string `json:"startup_on_boot"`
			CloseToTray              string `json:"close_to_tray"`
			DeeplAPIKey              string `json:"deepl_api_key"`
			DeeplEndpoint            string `json:"deepl_endpoint"`
			BaiduAppId               string `json:"baidu_app_id"`
			BaiduSecretKey           string `json:"baidu_secret_key"`
			MaxCacheSizeMb           string `json:"max_cache_size_mb"`
			DefaultViewMode          string `json:"default_view_mode"`
			ShowHiddenArticles       string `json:"show_hidden_articles"`
			SummaryEnabled           string `json:"summary_enabled"`
			MediaCacheMaxSizeMb      string `json:"media_cache_max_size_mb"`
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
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.MediaCacheMaxAgeDays != "" {
			h.DB.SetSetting("media_cache_max_age_days", req.MediaCacheMaxAgeDays)
		}

		if req.GoogleTranslateEndpoint != "" {
			h.DB.SetSetting("google_translate_endpoint", req.GoogleTranslateEndpoint)
		}

		if req.ShowArticlePreviewImages != "" {
			h.DB.SetSetting("show_article_preview_images", req.ShowArticlePreviewImages)
		}

		if req.NetworkLatencyMs != "" {
			h.DB.SetSetting("network_latency_ms", req.NetworkLatencyMs)
		}

		if req.SummaryLength != "" {
			h.DB.SetSetting("summary_length", req.SummaryLength)
		}

		if req.MaxArticleAgeDays != "" {
			h.DB.SetSetting("max_article_age_days", req.MaxArticleAgeDays)
		}

		if req.ProxyPort != "" {
			h.DB.SetSetting("proxy_port", req.ProxyPort)
		}

		if req.HoverMarkAsRead != "" {
			h.DB.SetSetting("hover_mark_as_read", req.HoverMarkAsRead)
		}

		if req.AIEndpoint != "" {
			h.DB.SetSetting("ai_endpoint", req.AIEndpoint)
		}

		if req.ObsidianVaultPath != "" {
			h.DB.SetSetting("obsidian_vault_path", req.ObsidianVaultPath)
		}

		if req.AIModel != "" {
			h.DB.SetSetting("ai_model", req.AIModel)
		}

		if req.WindowY != "" {
			h.DB.SetSetting("window_y", req.WindowY)
		}

		if req.AICustomHeaders != "" {
			h.DB.SetSetting("ai_custom_headers", req.AICustomHeaders)
		}

		if req.AutoCleanupEnabled != "" {
			h.DB.SetSetting("auto_cleanup_enabled", req.AutoCleanupEnabled)
		}

		if req.ObsidianEnabled != "" {
			h.DB.SetSetting("obsidian_enabled", req.ObsidianEnabled)
		}

		if req.SummaryProvider != "" {
			h.DB.SetSetting("summary_provider", req.SummaryProvider)
		}

		if err := h.DB.SetEncryptedSetting("freshrss_api_password", req.FreshRSSAPIPassword); err != nil {
			log.Printf("Failed to save freshrss_api_password: %v", err)
			http.Error(w, "Failed to save freshrss_api_password", http.StatusInternalServerError)
			return
		}

		if req.FullTextFetchEnabled != "" {
			h.DB.SetSetting("full_text_fetch_enabled", req.FullTextFetchEnabled)
		}

		if req.Theme != "" {
			h.DB.SetSetting("theme", req.Theme)
		}

		if req.AITranslationPrompt != "" {
			h.DB.SetSetting("ai_translation_prompt", req.AITranslationPrompt)
		}

		if req.ObsidianVault != "" {
			h.DB.SetSetting("obsidian_vault", req.ObsidianVault)
		}

		if req.Language != "" {
			h.DB.SetSetting("language", req.Language)
		}

		if err := h.DB.SetEncryptedSetting("ai_api_key", req.AIAPIKey); err != nil {
			log.Printf("Failed to save ai_api_key: %v", err)
			http.Error(w, "Failed to save ai_api_key", http.StatusInternalServerError)
			return
		}

		if req.MediaCacheEnabled != "" {
			h.DB.SetSetting("media_cache_enabled", req.MediaCacheEnabled)
		}

		if err := h.DB.SetEncryptedSetting("proxy_username", req.ProxyUsername); err != nil {
			log.Printf("Failed to save proxy_username: %v", err)
			http.Error(w, "Failed to save proxy_username", http.StatusInternalServerError)
			return
		}

		if req.MaxConcurrentRefreshes != "" {
			h.DB.SetSetting("max_concurrent_refreshes", req.MaxConcurrentRefreshes)
		}

		if req.AISummaryPrompt != "" {
			h.DB.SetSetting("ai_summary_prompt", req.AISummaryPrompt)
		}

		if err := h.DB.SetEncryptedSetting("proxy_password", req.ProxyPassword); err != nil {
			log.Printf("Failed to save proxy_password: %v", err)
			http.Error(w, "Failed to save proxy_password", http.StatusInternalServerError)
			return
		}

		if req.LastArticleUpdate != "" {
			h.DB.SetSetting("last_article_update", req.LastArticleUpdate)
		}

		if req.LastNetworkTest != "" {
			h.DB.SetSetting("last_network_test", req.LastNetworkTest)
		}

		if req.ImageGalleryEnabled != "" {
			h.DB.SetSetting("image_gallery_enabled", req.ImageGalleryEnabled)
		}

		if req.UpdateInterval != "" {
			h.DB.SetSetting("update_interval", req.UpdateInterval)
		}

		if req.AIUsageLimit != "" {
			h.DB.SetSetting("ai_usage_limit", req.AIUsageLimit)
		}

		if req.AIChatEnabled != "" {
			h.DB.SetSetting("ai_chat_enabled", req.AIChatEnabled)
		}

		if req.SummaryTriggerMode != "" {
			h.DB.SetSetting("summary_trigger_mode", req.SummaryTriggerMode)
		}

		if req.WindowHeight != "" {
			h.DB.SetSetting("window_height", req.WindowHeight)
		}

		if req.WindowMaximized != "" {
			h.DB.SetSetting("window_maximized", req.WindowMaximized)
		}

		if req.FreshRSSEnabled != "" {
			h.DB.SetSetting("freshrss_enabled", req.FreshRSSEnabled)
		}

		if req.RefreshMode != "" {
			h.DB.SetSetting("refresh_mode", req.RefreshMode)
		}

		if req.TranslationEnabled != "" {
			h.DB.SetSetting("translation_enabled", req.TranslationEnabled)
		}

		if req.TargetLanguage != "" {
			h.DB.SetSetting("target_language", req.TargetLanguage)
		}

		if req.ProxyEnabled != "" {
			h.DB.SetSetting("proxy_enabled", req.ProxyEnabled)
		}

		if req.AutoShowAllContent != "" {
			h.DB.SetSetting("auto_show_all_content", req.AutoShowAllContent)
		}

		if req.AIUsageTokens != "" {
			h.DB.SetSetting("ai_usage_tokens", req.AIUsageTokens)
		}

		if req.StartupOnBoot != "" {
			h.DB.SetSetting("startup_on_boot", req.StartupOnBoot)
		}

		if req.CloseToTray != "" {
			h.DB.SetSetting("close_to_tray", req.CloseToTray)
		}

		if err := h.DB.SetEncryptedSetting("deepl_api_key", req.DeeplAPIKey); err != nil {
			log.Printf("Failed to save deepl_api_key: %v", err)
			http.Error(w, "Failed to save deepl_api_key", http.StatusInternalServerError)
			return
		}

		if req.DeeplEndpoint != "" {
			h.DB.SetSetting("deepl_endpoint", req.DeeplEndpoint)
		}

		if req.BaiduAppId != "" {
			h.DB.SetSetting("baidu_app_id", req.BaiduAppId)
		}

		if err := h.DB.SetEncryptedSetting("baidu_secret_key", req.BaiduSecretKey); err != nil {
			log.Printf("Failed to save baidu_secret_key: %v", err)
			http.Error(w, "Failed to save baidu_secret_key", http.StatusInternalServerError)
			return
		}

		if req.MaxCacheSizeMb != "" {
			h.DB.SetSetting("max_cache_size_mb", req.MaxCacheSizeMb)
		}

		if req.DefaultViewMode != "" {
			h.DB.SetSetting("default_view_mode", req.DefaultViewMode)
		}

		if req.ShowHiddenArticles != "" {
			h.DB.SetSetting("show_hidden_articles", req.ShowHiddenArticles)
		}

		if req.SummaryEnabled != "" {
			h.DB.SetSetting("summary_enabled", req.SummaryEnabled)
		}

		if req.MediaCacheMaxSizeMb != "" {
			h.DB.SetSetting("media_cache_max_size_mb", req.MediaCacheMaxSizeMb)
		}

		if req.ProxyType != "" {
			h.DB.SetSetting("proxy_type", req.ProxyType)
		}

		if req.Rules != "" {
			h.DB.SetSetting("rules", req.Rules)
		}

		if req.WindowWidth != "" {
			h.DB.SetSetting("window_width", req.WindowWidth)
		}

		if req.NetworkSpeed != "" {
			h.DB.SetSetting("network_speed", req.NetworkSpeed)
		}

		if req.Shortcuts != "" {
			h.DB.SetSetting("shortcuts", req.Shortcuts)
		}

		if req.NetworkBandwidthMbps != "" {
			h.DB.SetSetting("network_bandwidth_mbps", req.NetworkBandwidthMbps)
		}

		if req.FreshRSSUsername != "" {
			h.DB.SetSetting("freshrss_username", req.FreshRSSUsername)
		}

		if req.TranslationProvider != "" {
			h.DB.SetSetting("translation_provider", req.TranslationProvider)
		}

		if req.ProxyHost != "" {
			h.DB.SetSetting("proxy_host", req.ProxyHost)
		}

		if req.WindowX != "" {
			h.DB.SetSetting("window_x", req.WindowX)
		}

		if req.FreshRSSServerUrl != "" {
			h.DB.SetSetting("freshrss_server_url", req.FreshRSSServerUrl)
		}
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
