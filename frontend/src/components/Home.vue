<script setup lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue'
import { Icon } from '@iconify/vue';
import RssList from './RssList.vue'
import RssListButton from './RssListButton.vue'
import RssContent from './RssContent.vue'
import RssContentButton from './RssContentButton.vue'
import { GetFeedContent, GetHistory, SetHistory, SetHistoryReaded, ClearHistory } from '../../wailsjs/go/main/App'

type FeedContent = {
  FeedTitle: string
  FeedImage: string
  Title: string
  Link: string
  TimeSince: string
  Time: string
  Image: string
  Content: string
  Readed: boolean
}

const feedContent = reactive({
  feedList: [] as FeedContent[],
})

const selectedFeed = ref<FeedContent | undefined>(undefined)

async function fetchFeedContent() {
  const result: FeedContent[] = await GetFeedContent()
  feedContent.feedList = result
  return feedContent
}

async function fetchHistoryContent() {
  const result: FeedContent[] = await GetHistory()
  feedContent.feedList = result
  return feedContent
}

async function setHistoryContent() {
  await SetHistory(feedContent.feedList)
}

async function deleteHistoryContent() {
  feedContent.feedList = []
  await ClearHistory()
  await handleClickRefresh()
}

const isRefreshing = ref(false)

async function handleClickRefresh() {
  isRefreshing.value = true
  await fetchHistoryContent()
  await fetchFeedContent()
  await setHistoryContent()
  isRefreshing.value = false
}

async function handleFeedClicked(feed: FeedContent) {
  selectedFeed.value = feed
  await modifyFeedContentReaded(feed, true)
}

async function modifyFeedContentReaded(feed: FeedContent, readed: boolean) {
  const index = feedContent.feedList.findIndex((f) => f.Link === feed.Link)
  if (index !== -1) {
    if (feedContent.feedList[index].Readed !== readed) {
      feedContent.feedList[index].Readed = readed
      await SetHistoryReaded(feedContent.feedList[index])
    }
  }
}

defineComponent({
  components: {
    feedContent,
  },
  setup(_, { emit }) {
    return {
      RssContentButton,
      isRefreshing,
      handleClickRefresh,
      deleteHistoryContent,
      modifyFeedContentReaded,
    }
  }
})

onMounted(async () => {
  await handleClickRefresh()
})
</script>

<template>
  <aside>
    <rss-list-button 
      @delete-history-content="deleteHistoryContent" 
      @handle-click-refresh="handleClickRefresh" 
      :isRefreshing="isRefreshing"
    />
    <rss-list 
      @feed-clicked="handleFeedClicked" 
      :feedContent="feedContent" 
    />
  </aside>
  <main>
    <rss-content-button 
      v-if="selectedFeed !== undefined"
      @modify-feed-content-readed="modifyFeedContentReaded" 
      :selectedFeed="selectedFeed"
    />
    <rss-content 
      v-if="selectedFeed" 
      :selectedFeed="selectedFeed" 
    />
    <div v-else class="NoSelectedFeed"></div>
  </main>
</template>

<style>
aside {
  display: flex;
  flex-direction: column;

  min-width: 344px;
  max-width: 344px;
  height: 100vh;

  color: #000000;
  background-color: #f0f0f0;

  word-wrap: normal;
}

main {
  display: flex;
  flex-direction: column;

  width: calc(100vw - 344px);

  border-left: 1px solid #ccc;

  height: 100vh;
}

.NoSelectedFeed {
  width: 100%;
  height: 100%;
  background-color: #f0f0f0;
}
</style>