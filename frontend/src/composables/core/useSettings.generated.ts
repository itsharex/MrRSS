// Copyright 2026 Ch3nyang & MrRSS Team. All rights reserved.
//
// Auto-generated settings composable functions
// CODE GENERATED - DO NOT EDIT MANUALLY
// To add new settings, edit internal/config/settings_schema.json and run: go run tools/settings-generator/main.go
import { type Ref } from 'vue';
import type { SettingsData } from '@/types/settings.generated';

// Generated settings defaults
export const settingsDefaults = {
  summaryProvider: 'local',
  freshRSSAPIPassword: '',
  fullTextFetchEnabled: true,
  theme: 'auto',
  aiTranslationPrompt:
    'You are a translator. Translate the given text accurately. Output ONLY the translated text, nothing else.',
  obsidianVault: '',
  language: 'en-US',
  aiAPIKey: '',
  mediaCacheEnabled: false,
  proxyUsername: '',
  maxConcurrentRefreshes: '5',
  aiSummaryPrompt:
    'You are a summarizer. Generate a concise summary of the given text. Output ONLY the summary, nothing else.',
  proxyPassword: '',
  lastArticleUpdate: '',
  lastNetworkTest: '',
  imageGalleryEnabled: false,
  updateInterval: 30,
  aiUsageLimit: '200',
  aiChatEnabled: false,
  summaryTriggerMode: 'manual',
  windowHeight: '768',
  windowMaximized: 'false',
  freshRSSSyncEnabled: false,
  refreshMode: 'fixed',
  translationEnabled: false,
  targetLanguage: 'zh',
  proxyEnabled: false,
  autoShowAllContent: false,
  aiUsageTokens: '0',
  startupOnBoot: false,
  closeToTray: true,
  deeplAPIKey: '',
  deeplEndpoint: '',
  baiduAppID: '',
  baiduSecretKey: '',
  maxCacheSizeMB: 20,
  defaultViewMode: 'rendered',
  showHiddenArticles: false,
  summaryEnabled: true,
  mediaCacheMaxSizeMB: 100,
  proxyType: 'https',
  rules: '',
  windowWidth: '1024',
  networkSpeed: 'medium',
  shortcuts: '',
  networkBandwidth: '0',
  freshRSSUsername: '',
  translationProvider: 'google',
  proxyHost: '127.0.0.1',
  windowX: '0',
  freshRSSServerURL: '',
  mediaCacheMaxAgeDays: 7,
  googleTranslateEndpoint: 'translate.googleapis.com',
  showArticlePreviewImages: true,
  networkLatency: '0',
  summaryLength: 'medium',
  maxArticleAgeDays: 30,
  proxyPort: '7890',
  hoverMarkAsRead: false,
  aiEndpoint: 'https://api.openai.com/v1/chat/completions',
  obsidianVaultPath: '',
  aiModel: 'gpt-4o-mini',
  windowY: '0',
  aiCustomHeaders: '',
  autoCleanupEnabled: false,
  obsidianEnabled: false,
} as const;

// Generated fetchSettings function
export function useGeneratedFetchSettings(settingsRef: Ref<SettingsData>) {
  return async () => {
    const response = await fetch('/api/settings');
    if (!response.ok) {
      throw new Error('Failed to fetch settings');
    }
    const data = await response.json();
    settingsRef.value = {
      aiSummaryPrompt:
        data.ai_summary_prompt ||
        'You are a summarizer. Generate a concise summary of the given text. Output ONLY the summary, nothing else.',
      proxyPassword: data.proxy_password || '',
      lastArticleUpdate: data.last_article_update || '',
      lastNetworkTest: data.last_network_test || '',
      imageGalleryEnabled: data.image_gallery_enabled === 'true',
      updateInterval: parseInt(data.update_interval) || 30,
      aiUsageLimit: data.ai_usage_limit || '200',
      aiChatEnabled: data.ai_chat_enabled === 'true',
      summaryTriggerMode: data.summary_trigger_mode || 'manual',
      windowHeight: data.window_height || '768',
      windowMaximized: data.window_maximized || 'false',
      freshRSSSyncEnabled: data.freshrss_sync_enabled === 'true',
      refreshMode: data.refresh_mode || 'fixed',
      translationEnabled: data.translation_enabled === 'true',
      targetLanguage: data.target_language || 'zh',
      proxyEnabled: data.proxy_enabled === 'true',
      autoShowAllContent: data.auto_show_all_content === 'true',
      aiUsageTokens: data.ai_usage_tokens || '0',
      startupOnBoot: data.startup_on_boot === 'true',
      closeToTray: data.close_to_tray === 'true',
      deeplAPIKey: data.deepl_a_p_i_key || '',
      deeplEndpoint: data.deepl_endpoint || '',
      baiduAppID: data.baidu_app_i_d || '',
      baiduSecretKey: data.baidu_secret_key || '',
      maxCacheSizeMB: parseInt(data.max_cache_size_m_b) || 20,
      defaultViewMode: data.default_view_mode || 'rendered',
      showHiddenArticles: data.show_hidden_articles === 'true',
      summaryEnabled: data.summary_enabled === 'true',
      mediaCacheMaxSizeMB: parseInt(data.media_cache_max_size_m_b) || 100,
      proxyType: data.proxy_type || 'https',
      rules: data.rules || '',
      windowWidth: data.window_width || '1024',
      networkSpeed: data.network_speed || 'medium',
      shortcuts: data.shortcuts || '',
      networkBandwidth: data.network_bandwidth || '0',
      freshRSSUsername: data.freshrss_username || '',
      translationProvider: data.translation_provider || 'google',
      proxyHost: data.proxy_host || '127.0.0.1',
      windowX: data.window_x || '0',
      freshRSSServerURL: data.freshrss_server_u_r_l || '',
      mediaCacheMaxAgeDays: parseInt(data.media_cache_max_age_days) || 7,
      googleTranslateEndpoint: data.google_translate_endpoint || 'translate.googleapis.com',
      showArticlePreviewImages: data.show_article_preview_images === 'true',
      networkLatency: data.network_latency || '0',
      summaryLength: data.summary_length || 'medium',
      maxArticleAgeDays: parseInt(data.max_article_age_days) || 30,
      proxyPort: data.proxy_port || '7890',
      hoverMarkAsRead: data.hover_mark_as_read === 'true',
      aiEndpoint: data.ai_endpoint || 'https://api.openai.com/v1/chat/completions',
      obsidianVaultPath: data.obsidian_vault_path || '',
      aiModel: data.ai_model || 'gpt-4o-mini',
      windowY: data.window_y || '0',
      aiCustomHeaders: data.ai_custom_headers || '',
      autoCleanupEnabled: data.auto_cleanup_enabled === 'true',
      obsidianEnabled: data.obsidian_enabled === 'true',
      summaryProvider: data.summary_provider || 'local',
      freshRSSAPIPassword: data.freshrss_a_p_i_password || '',
      fullTextFetchEnabled: data.full_text_fetch_enabled === 'true',
      theme: data.theme || 'auto',
      aiTranslationPrompt:
        data.ai_translation_prompt ||
        'You are a translator. Translate the given text accurately. Output ONLY the translated text, nothing else.',
      obsidianVault: data.obsidian_vault || '',
      language: data.language || 'en-US',
      aiAPIKey: data.ai_a_p_i_key || '',
      mediaCacheEnabled: data.media_cache_enabled === 'true',
      proxyUsername: data.proxy_username || '',
      maxConcurrentRefreshes: data.max_concurrent_refreshes || '5',
    };
  };
}

// Generated auto-save payload builder
export function buildSavePayload(settingsRef: Ref<SettingsData>) {
  return {
    ai_summary_prompt: settingsRef.value.aiSummaryPrompt ?? settingsDefaults.aiSummaryPrompt,
    proxy_password: settingsRef.value.proxyPassword ?? settingsDefaults.proxyPassword,
    last_article_update: settingsRef.value.lastArticleUpdate ?? settingsDefaults.lastArticleUpdate,
    last_network_test: settingsRef.value.lastNetworkTest ?? settingsDefaults.lastNetworkTest,
    image_gallery_enabled: (
      settingsRef.value.imageGalleryEnabled ?? settingsDefaults.imageGalleryEnabled
    ).toString(),
    update_interval: settingsRef.value.updateInterval ?? settingsDefaults.updateInterval,
    ai_usage_limit: settingsRef.value.aiUsageLimit ?? settingsDefaults.aiUsageLimit,
    ai_chat_enabled: (settingsRef.value.aiChatEnabled ?? settingsDefaults.aiChatEnabled).toString(),
    summary_trigger_mode:
      settingsRef.value.summaryTriggerMode ?? settingsDefaults.summaryTriggerMode,
    window_height: settingsRef.value.windowHeight ?? settingsDefaults.windowHeight,
    window_maximized: settingsRef.value.windowMaximized ?? settingsDefaults.windowMaximized,
    freshrss_sync_enabled: (
      settingsRef.value.freshRSSSyncEnabled ?? settingsDefaults.freshRSSSyncEnabled
    ).toString(),
    refresh_mode: settingsRef.value.refreshMode ?? settingsDefaults.refreshMode,
    translation_enabled: (
      settingsRef.value.translationEnabled ?? settingsDefaults.translationEnabled
    ).toString(),
    target_language: settingsRef.value.targetLanguage ?? settingsDefaults.targetLanguage,
    proxy_enabled: (settingsRef.value.proxyEnabled ?? settingsDefaults.proxyEnabled).toString(),
    auto_show_all_content: (
      settingsRef.value.autoShowAllContent ?? settingsDefaults.autoShowAllContent
    ).toString(),
    ai_usage_tokens: settingsRef.value.aiUsageTokens ?? settingsDefaults.aiUsageTokens,
    startup_on_boot: (settingsRef.value.startupOnBoot ?? settingsDefaults.startupOnBoot).toString(),
    close_to_tray: (settingsRef.value.closeToTray ?? settingsDefaults.closeToTray).toString(),
    deepl_a_p_i_key: settingsRef.value.deeplAPIKey ?? settingsDefaults.deeplAPIKey,
    deepl_endpoint: settingsRef.value.deeplEndpoint ?? settingsDefaults.deeplEndpoint,
    baidu_app_i_d: settingsRef.value.baiduAppID ?? settingsDefaults.baiduAppID,
    baidu_secret_key: settingsRef.value.baiduSecretKey ?? settingsDefaults.baiduSecretKey,
    max_cache_size_m_b: settingsRef.value.maxCacheSizeMB ?? settingsDefaults.maxCacheSizeMB,
    default_view_mode: settingsRef.value.defaultViewMode ?? settingsDefaults.defaultViewMode,
    show_hidden_articles: (
      settingsRef.value.showHiddenArticles ?? settingsDefaults.showHiddenArticles
    ).toString(),
    summary_enabled: (
      settingsRef.value.summaryEnabled ?? settingsDefaults.summaryEnabled
    ).toString(),
    media_cache_max_size_m_b:
      settingsRef.value.mediaCacheMaxSizeMB ?? settingsDefaults.mediaCacheMaxSizeMB,
    proxy_type: settingsRef.value.proxyType ?? settingsDefaults.proxyType,
    rules: settingsRef.value.rules ?? settingsDefaults.rules,
    window_width: settingsRef.value.windowWidth ?? settingsDefaults.windowWidth,
    network_speed: settingsRef.value.networkSpeed ?? settingsDefaults.networkSpeed,
    shortcuts: settingsRef.value.shortcuts ?? settingsDefaults.shortcuts,
    network_bandwidth: settingsRef.value.networkBandwidth ?? settingsDefaults.networkBandwidth,
    freshrss_username: settingsRef.value.freshRSSUsername ?? settingsDefaults.freshRSSUsername,
    translation_provider:
      settingsRef.value.translationProvider ?? settingsDefaults.translationProvider,
    proxy_host: settingsRef.value.proxyHost ?? settingsDefaults.proxyHost,
    window_x: settingsRef.value.windowX ?? settingsDefaults.windowX,
    freshrss_server_u_r_l:
      settingsRef.value.freshRSSServerURL ?? settingsDefaults.freshRSSServerURL,
    media_cache_max_age_days:
      settingsRef.value.mediaCacheMaxAgeDays ?? settingsDefaults.mediaCacheMaxAgeDays,
    google_translate_endpoint:
      settingsRef.value.googleTranslateEndpoint ?? settingsDefaults.googleTranslateEndpoint,
    show_article_preview_images: (
      settingsRef.value.showArticlePreviewImages ?? settingsDefaults.showArticlePreviewImages
    ).toString(),
    network_latency: settingsRef.value.networkLatency ?? settingsDefaults.networkLatency,
    summary_length: settingsRef.value.summaryLength ?? settingsDefaults.summaryLength,
    max_article_age_days: settingsRef.value.maxArticleAgeDays ?? settingsDefaults.maxArticleAgeDays,
    proxy_port: settingsRef.value.proxyPort ?? settingsDefaults.proxyPort,
    hover_mark_as_read: (
      settingsRef.value.hoverMarkAsRead ?? settingsDefaults.hoverMarkAsRead
    ).toString(),
    ai_endpoint: settingsRef.value.aiEndpoint ?? settingsDefaults.aiEndpoint,
    obsidian_vault_path: settingsRef.value.obsidianVaultPath ?? settingsDefaults.obsidianVaultPath,
    ai_model: settingsRef.value.aiModel ?? settingsDefaults.aiModel,
    window_y: settingsRef.value.windowY ?? settingsDefaults.windowY,
    ai_custom_headers: settingsRef.value.aiCustomHeaders ?? settingsDefaults.aiCustomHeaders,
    auto_cleanup_enabled: (
      settingsRef.value.autoCleanupEnabled ?? settingsDefaults.autoCleanupEnabled
    ).toString(),
    obsidian_enabled: (
      settingsRef.value.obsidianEnabled ?? settingsDefaults.obsidianEnabled
    ).toString(),
    summary_provider: settingsRef.value.summaryProvider ?? settingsDefaults.summaryProvider,
    freshrss_a_p_i_password:
      settingsRef.value.freshRSSAPIPassword ?? settingsDefaults.freshRSSAPIPassword,
    full_text_fetch_enabled: (
      settingsRef.value.fullTextFetchEnabled ?? settingsDefaults.fullTextFetchEnabled
    ).toString(),
    theme: settingsRef.value.theme ?? settingsDefaults.theme,
    ai_translation_prompt:
      settingsRef.value.aiTranslationPrompt ?? settingsDefaults.aiTranslationPrompt,
    obsidian_vault: settingsRef.value.obsidianVault ?? settingsDefaults.obsidianVault,
    language: settingsRef.value.language ?? settingsDefaults.language,
    ai_a_p_i_key: settingsRef.value.aiAPIKey ?? settingsDefaults.aiAPIKey,
    media_cache_enabled: (
      settingsRef.value.mediaCacheEnabled ?? settingsDefaults.mediaCacheEnabled
    ).toString(),
    proxy_username: settingsRef.value.proxyUsername ?? settingsDefaults.proxyUsername,
    max_concurrent_refreshes:
      settingsRef.value.maxConcurrentRefreshes ?? settingsDefaults.maxConcurrentRefreshes,
  };
}

// Generated event dispatchers
export function dispatchSettingChangeEvents(settingsRef: Ref<SettingsData>) {
  window.dispatchEvent(
    new CustomEvent('ai-summary-prompt-changed', {
      detail: { value: settingsRef.value.aiSummaryPrompt },
    })
  );
  window.dispatchEvent(
    new CustomEvent('proxy-password-changed', {
      detail: { value: settingsRef.value.proxyPassword },
    })
  );
  window.dispatchEvent(
    new CustomEvent('last-network-test-changed', {
      detail: { value: settingsRef.value.lastNetworkTest },
    })
  );
  window.dispatchEvent(
    new CustomEvent('image-gallery-enabled-changed', {
      detail: { value: settingsRef.value.imageGalleryEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('update-interval-changed', {
      detail: { value: settingsRef.value.updateInterval },
    })
  );
  window.dispatchEvent(
    new CustomEvent('ai-usage-limit-changed', { detail: { value: settingsRef.value.aiUsageLimit } })
  );
  window.dispatchEvent(
    new CustomEvent('ai-chat-enabled-changed', {
      detail: { value: settingsRef.value.aiChatEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('summary-trigger-mode-changed', {
      detail: { value: settingsRef.value.summaryTriggerMode },
    })
  );
  window.dispatchEvent(
    new CustomEvent('fresh-r-s-s-sync-enabled-changed', {
      detail: { value: settingsRef.value.freshRSSSyncEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('refresh-mode-changed', { detail: { value: settingsRef.value.refreshMode } })
  );
  window.dispatchEvent(
    new CustomEvent('translation-enabled-changed', {
      detail: { value: settingsRef.value.translationEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('target-language-changed', {
      detail: { value: settingsRef.value.targetLanguage },
    })
  );
  window.dispatchEvent(
    new CustomEvent('proxy-enabled-changed', { detail: { value: settingsRef.value.proxyEnabled } })
  );
  window.dispatchEvent(
    new CustomEvent('auto-show-all-content-changed', {
      detail: { value: settingsRef.value.autoShowAllContent },
    })
  );
  window.dispatchEvent(
    new CustomEvent('ai-usage-tokens-changed', {
      detail: { value: settingsRef.value.aiUsageTokens },
    })
  );
  window.dispatchEvent(
    new CustomEvent('startup-on-boot-changed', {
      detail: { value: settingsRef.value.startupOnBoot },
    })
  );
  window.dispatchEvent(
    new CustomEvent('close-to-tray-changed', { detail: { value: settingsRef.value.closeToTray } })
  );
  window.dispatchEvent(
    new CustomEvent('deepl-a-p-i-key-changed', { detail: { value: settingsRef.value.deeplAPIKey } })
  );
  window.dispatchEvent(
    new CustomEvent('deepl-endpoint-changed', {
      detail: { value: settingsRef.value.deeplEndpoint },
    })
  );
  window.dispatchEvent(
    new CustomEvent('baidu-app-i-d-changed', { detail: { value: settingsRef.value.baiduAppID } })
  );
  window.dispatchEvent(
    new CustomEvent('baidu-secret-key-changed', {
      detail: { value: settingsRef.value.baiduSecretKey },
    })
  );
  window.dispatchEvent(
    new CustomEvent('max-cache-size-m-b-changed', {
      detail: { value: settingsRef.value.maxCacheSizeMB },
    })
  );
  window.dispatchEvent(
    new CustomEvent('default-view-mode-changed', {
      detail: { value: settingsRef.value.defaultViewMode },
    })
  );
  window.dispatchEvent(
    new CustomEvent('show-hidden-articles-changed', {
      detail: { value: settingsRef.value.showHiddenArticles },
    })
  );
  window.dispatchEvent(
    new CustomEvent('summary-enabled-changed', {
      detail: { value: settingsRef.value.summaryEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('media-cache-max-size-m-b-changed', {
      detail: { value: settingsRef.value.mediaCacheMaxSizeMB },
    })
  );
  window.dispatchEvent(
    new CustomEvent('proxy-type-changed', { detail: { value: settingsRef.value.proxyType } })
  );
  window.dispatchEvent(
    new CustomEvent('rules-changed', { detail: { value: settingsRef.value.rules } })
  );
  window.dispatchEvent(
    new CustomEvent('network-speed-changed', { detail: { value: settingsRef.value.networkSpeed } })
  );
  window.dispatchEvent(
    new CustomEvent('shortcuts-changed', { detail: { value: settingsRef.value.shortcuts } })
  );
  window.dispatchEvent(
    new CustomEvent('network-bandwidth-changed', {
      detail: { value: settingsRef.value.networkBandwidth },
    })
  );
  window.dispatchEvent(
    new CustomEvent('fresh-r-s-s-username-changed', {
      detail: { value: settingsRef.value.freshRSSUsername },
    })
  );
  window.dispatchEvent(
    new CustomEvent('translation-provider-changed', {
      detail: { value: settingsRef.value.translationProvider },
    })
  );
  window.dispatchEvent(
    new CustomEvent('proxy-host-changed', { detail: { value: settingsRef.value.proxyHost } })
  );
  window.dispatchEvent(
    new CustomEvent('fresh-r-s-s-server-u-r-l-changed', {
      detail: { value: settingsRef.value.freshRSSServerURL },
    })
  );
  window.dispatchEvent(
    new CustomEvent('media-cache-max-age-days-changed', {
      detail: { value: settingsRef.value.mediaCacheMaxAgeDays },
    })
  );
  window.dispatchEvent(
    new CustomEvent('google-translate-endpoint-changed', {
      detail: { value: settingsRef.value.googleTranslateEndpoint },
    })
  );
  window.dispatchEvent(
    new CustomEvent('show-article-preview-images-changed', {
      detail: { value: settingsRef.value.showArticlePreviewImages },
    })
  );
  window.dispatchEvent(
    new CustomEvent('network-latency-changed', {
      detail: { value: settingsRef.value.networkLatency },
    })
  );
  window.dispatchEvent(
    new CustomEvent('summary-length-changed', {
      detail: { value: settingsRef.value.summaryLength },
    })
  );
  window.dispatchEvent(
    new CustomEvent('max-article-age-days-changed', {
      detail: { value: settingsRef.value.maxArticleAgeDays },
    })
  );
  window.dispatchEvent(
    new CustomEvent('proxy-port-changed', { detail: { value: settingsRef.value.proxyPort } })
  );
  window.dispatchEvent(
    new CustomEvent('hover-mark-as-read-changed', {
      detail: { value: settingsRef.value.hoverMarkAsRead },
    })
  );
  window.dispatchEvent(
    new CustomEvent('ai-endpoint-changed', { detail: { value: settingsRef.value.aiEndpoint } })
  );
  window.dispatchEvent(
    new CustomEvent('obsidian-vault-path-changed', {
      detail: { value: settingsRef.value.obsidianVaultPath },
    })
  );
  window.dispatchEvent(
    new CustomEvent('ai-model-changed', { detail: { value: settingsRef.value.aiModel } })
  );
  window.dispatchEvent(
    new CustomEvent('ai-custom-headers-changed', {
      detail: { value: settingsRef.value.aiCustomHeaders },
    })
  );
  window.dispatchEvent(
    new CustomEvent('auto-cleanup-enabled-changed', {
      detail: { value: settingsRef.value.autoCleanupEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('obsidian-enabled-changed', {
      detail: { value: settingsRef.value.obsidianEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('summary-provider-changed', {
      detail: { value: settingsRef.value.summaryProvider },
    })
  );
  window.dispatchEvent(
    new CustomEvent('fresh-r-s-s-a-p-i-password-changed', {
      detail: { value: settingsRef.value.freshRSSAPIPassword },
    })
  );
  window.dispatchEvent(
    new CustomEvent('full-text-fetch-enabled-changed', {
      detail: { value: settingsRef.value.fullTextFetchEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('theme-changed', { detail: { value: settingsRef.value.theme } })
  );
  window.dispatchEvent(
    new CustomEvent('ai-translation-prompt-changed', {
      detail: { value: settingsRef.value.aiTranslationPrompt },
    })
  );
  window.dispatchEvent(
    new CustomEvent('obsidian-vault-changed', {
      detail: { value: settingsRef.value.obsidianVault },
    })
  );
  window.dispatchEvent(
    new CustomEvent('language-changed', { detail: { value: settingsRef.value.language } })
  );
  window.dispatchEvent(
    new CustomEvent('ai-a-p-i-key-changed', { detail: { value: settingsRef.value.aiAPIKey } })
  );
  window.dispatchEvent(
    new CustomEvent('media-cache-enabled-changed', {
      detail: { value: settingsRef.value.mediaCacheEnabled },
    })
  );
  window.dispatchEvent(
    new CustomEvent('proxy-username-changed', {
      detail: { value: settingsRef.value.proxyUsername },
    })
  );
  window.dispatchEvent(
    new CustomEvent('max-concurrent-refreshes-changed', {
      detail: { value: settingsRef.value.maxConcurrentRefreshes },
    })
  );
}
