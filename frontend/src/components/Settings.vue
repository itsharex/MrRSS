<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ref, watchEffect, onMounted, computed } from 'vue';
import { Icon } from '@iconify/vue';

import { GetFeedList, SetFeedList, DeleteFeedList } from '../../wailsjs/go/main/App'

type FeedList = {
  Link: string
  Category: string
}

const feedList = ref([] as FeedList[])

async function getFeedList() {
  const result: FeedList[] = await GetFeedList()
  console.log(result)
  feedList.value = result
}

const selectedSubscribeType = ref('RSS/Atom');
const subscribeUrl = ref('');

async function setFeedList() {
  const feed: FeedList = {
    Link: subscribeUrl.value,
    Category: selectedSubscribeType.value
  }

  console.log(feed)
  console.log([feed] as FeedList[])

  await SetFeedList([feed] as FeedList[])
  await getFeedList()

  selectedSubscribeType.value = 'RSS/Atom'
  subscribeUrl.value = ''
}

async function deleteFeedList(feed: FeedList) {
  await DeleteFeedList(feed)
  await getFeedList()
}

const router = useRouter();

const goBack = () => {
  router.push('/');
}

let subscribeUrlLabel = ref('');

watchEffect(() => {
  switch (selectedSubscribeType.value) {
    case 'RSS/Atom':
      subscribeUrlLabel.value = 'URL';
      break;
    case 'Twitter':
      subscribeUrlLabel.value = 'Username';
      break;
    case 'Telegram':
      subscribeUrlLabel.value = 'ID';
      break;
    case 'Youtube':
      subscribeUrlLabel.value = 'Username';
      break;
    case 'Wechat':
      subscribeUrlLabel.value = 'ID';
      break;
    default:
      subscribeUrlLabel.value = 'URL';
  }
});

const pageTitle = computed(() => {
  switch (router.currentRoute.value.path) {
    case '/settings/rss':
      return 'RSS Settings';
    case '/settings/preference':
      return 'Preferences';
    case '/settings/about':
      return 'About';
    default:
      return 'Settings';
  }
});

onMounted(() => {
  getFeedList()
})
</script>

<template>
  <div class="settings">
    <div class="nav">
      <div class="title">
        <button @click="goBack" class="btn done" title="Back to home">
          <Icon icon="material-symbols:arrow-back" />
        </button>
        <h1>{{ pageTitle }}</h1>
      </div>
      <div class="router">
        <router-link to="/settings/rss" class="btn" title="RSS settings">
          <Icon icon="material-symbols:rss-feed" />
        </router-link>
        <router-link to="/settings/preference" class="btn" title="Preference settings">
          <Icon icon="material-symbols:inbox-customize" />
        </router-link>
        <router-link to="/settings/about" class="btn" title="About MrRSS">
          <Icon icon="material-symbols:page-info" />
        </router-link>
      </div>
    </div>
    <router-view />
  </div>
</template>

<style scoped>
.btn {
  width: 32px;
  height: 32px;
  font-size: large;
  color: black;
  background-color: #e0e0e0;
  border: none;
  cursor: pointer;

  display: flex;
  justify-content: center;
  align-items: center;
}

.btn:hover {
  background-color: #d0d0d0;
}

.settings {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 32px;
  background-color: #f0f0f0;
  width: 100%;
  min-height: calc(100vh - 64px);
  margin: 0 auto;
  padding: 32px 0;
}

.settings>* {
  min-width: 560px;
  max-width: 560px;
  margin: 0 auto;
}

.nav {
  height: 48px;
  text-align: left;
  align-self: flex-start;

  display: flex;
  justify-content: space-between;
  align-items: center;
}

.title {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}

.router {
  display: flex;
  justify-content: center;
  align-items: center;
}

.router .btn {
  width: 48px;
  height: 46px;
  font-size: 24px;
}

.router .router-link-active {
  border-bottom: 2px solid #000;
}

/* 其它按钮 */
.router>:not(.router-link-active) {
  border-bottom: 2px solid #e0e0e0;
}

.done {
  height: 48px;
  min-width: 48px;
  border-radius: 50%;
  font-size: 32px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  margin: 0 0 0 auto;
}
</style>