<template>
  <div class="space-y-4">
    <h3 class="text-lg font-semibold text-text-primary">{{ t('networkSettings') }}</h3>
    <p class="text-sm text-text-secondary">{{ t('networkSettingsDescription') }}</p>

    <!-- Network Status Display -->
    <div class="bg-bg-secondary rounded-lg p-4 space-y-3">
      <!-- Speed Level -->
      <div class="flex items-center justify-between">
        <span class="text-sm text-text-secondary">{{ t('networkSpeed') }}</span>
        <div class="flex items-center gap-2">
          <span
            class="px-3 py-1 rounded-full text-xs font-medium"
            :class="speedLevelClass"
          >
            {{ t(`networkSpeed_${networkInfo.speed_level}`) }}
          </span>
        </div>
      </div>

      <!-- Bandwidth -->
      <div class="flex items-center justify-between">
        <span class="text-sm text-text-secondary">{{ t('bandwidth') }}</span>
        <span class="text-sm font-medium text-text-primary">
          {{ networkInfo.bandwidth_mbps.toFixed(2) }} Mbps
        </span>
      </div>

      <!-- Latency -->
      <div class="flex items-center justify-between">
        <span class="text-sm text-text-secondary">{{ t('latency') }}</span>
        <span class="text-sm font-medium text-text-primary">
          {{ networkInfo.latency_ms }} ms
        </span>
      </div>

      <!-- Max Concurrent Refreshes -->
      <div class="flex items-center justify-between">
        <span class="text-sm text-text-secondary">{{ t('maxConcurrentRefreshes') }}</span>
        <span class="text-sm font-medium text-text-primary">
          {{ networkInfo.max_concurrency }}
        </span>
      </div>

      <!-- Last Test Time -->
      <div v-if="networkInfo.detection_time" class="flex items-center justify-between">
        <span class="text-sm text-text-secondary">{{ t('lastDetection') }}</span>
        <span class="text-xs text-text-tertiary">
          {{ formatTime(networkInfo.detection_time) }}
        </span>
      </div>
    </div>

    <!-- Re-detect Button -->
    <button
      class="btn-primary w-full flex items-center justify-center gap-2"
      :disabled="isDetecting"
      @click="detectNetwork"
    >
      <i class="ph ph-arrow-clockwise" :class="{ 'animate-spin': isDetecting }"></i>
      <span>{{ isDetecting ? t('detecting') : t('reDetectNetwork') }}</span>
    </button>

    <!-- Error Message -->
    <div
      v-if="errorMessage"
      class="bg-red-500/10 border border-red-500/20 rounded-lg p-3 text-sm text-red-500"
    >
      {{ errorMessage }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import type { NetworkInfo } from '@/types/settings';

const { t } = useI18n();

const networkInfo = ref<NetworkInfo>({
  speed_level: 'medium',
  bandwidth_mbps: 0,
  latency_ms: 0,
  max_concurrency: 5,
  detection_time: '',
  detection_success: false,
});

const isDetecting = ref(false);
const errorMessage = ref('');

const speedLevelClass = computed(() => {
  switch (networkInfo.value.speed_level) {
    case 'fast':
      return 'bg-green-500/20 text-green-600 dark:text-green-400';
    case 'medium':
      return 'bg-yellow-500/20 text-yellow-600 dark:text-yellow-400';
    case 'slow':
      return 'bg-red-500/20 text-red-600 dark:text-red-400';
    default:
      return 'bg-gray-500/20 text-gray-600 dark:text-gray-400';
  }
});

async function loadNetworkInfo() {
  try {
    const response = await fetch('/api/network/info');
    if (response.ok) {
      const data = await response.json();
      networkInfo.value = data;
    }
  } catch (error) {
    console.error('Failed to load network info:', error);
  }
}

async function detectNetwork() {
  isDetecting.value = true;
  errorMessage.value = '';

  try {
    const response = await fetch('/api/network/detect', {
      method: 'POST',
    });

    if (response.ok) {
      const data = await response.json();
      networkInfo.value = data;

      if (!data.detection_success) {
        errorMessage.value = data.error_message || t('networkDetectionFailed');
      } else {
        window.showToast(t('networkDetectionComplete'), 'success');
      }
    } else {
      errorMessage.value = t('networkDetectionFailed');
    }
  } catch (error) {
    console.error('Network detection error:', error);
    errorMessage.value = t('networkDetectionFailed');
  } finally {
    isDetecting.value = false;
  }
}

function formatTime(timeStr: string): string {
  if (!timeStr) return '';
  const date = new Date(timeStr);
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  const minutes = Math.floor(diff / 60000);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  if (days > 0) {
    return t('daysAgo', { count: days });
  } else if (hours > 0) {
    return t('hoursAgo', { count: hours });
  } else if (minutes > 0) {
    return t('minutesAgo', { count: minutes });
  } else {
    return t('justNow');
  }
}

onMounted(() => {
  loadNetworkInfo();
});
</script>
