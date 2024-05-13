<script setup lang="ts">
import { ref, watchEffect, onMounted } from 'vue';
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

onMounted(() => {
  getFeedList()
})
</script>

<template>
  <form name="new feed">
    <label for="subscribe-type">RSS Type</label>
    <select id="subscribe-type" name="subscribe-type" v-model="selectedSubscribeType">
      <option value="RSS/Atom" selected>RSS/Atom</option>
      <option value="Twitter">Twitter</option>
      <option value="Telegram">Telegram</option>
      <option value="Youtube">Youtube</option>
      <option value="Wechat">Wechat Public Account</option>
    </select>
    <label for="subscribe-url">{{ subscribeUrlLabel }}</label>
    <input type="text" id="subscribe-url" name="subscribe-url" v-model="subscribeUrl" autocomplete="off"
      placeholder="https://feeds.bbci.co.uk/news/world/rss.xml" required />
    <button @click.prevent="setFeedList" class="btn" title="Add feed">
      <Icon icon="material-symbols:forms-add-on" />
    </button>
  </form>
  <ul>
    <li v-for="feed in feedList" :key="feed.Link">
      <div class="img">
        <img :src="`https://www.google.com/s2/favicons?domain=${feed.Link}`" alt="favicon" />
      </div>
      <span class="link">{{ feed.Link }}</span>
      <span class="category">{{ feed.Category }}</span>
      <button @click="deleteFeedList(feed)" class="btn" title="Delete feed">
        <Icon icon="material-symbols:delete-forever" />
      </button>
    </li>
  </ul>
</template>

<style scoped>
.btn {
  width: 32px;
  height: 32px;
  font-size: large;
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

form {
  width: 100%;

  display: grid;
  gap: 8px;
  grid-template-columns: 80px 1fr auto;
  grid-template-rows: 1fr 1fr;
}

form label {
  grid-column: 1 / 2;
  text-align: left;
  font-weight: bold;
  line-height: 24px;
}

form select,
form input {
  grid-column: 2 / 3;
  background-color: #f9f9f9;
  border: none;
}

form select:focus,
form input:focus {
  outline: none;
}

form .btn {
  grid-column: 3 / 4;
  grid-row: 1 / 3;
  height: 56px;
}

ul {
  list-style-type: none;
  padding: 0;
  width: 100%;
  max-height: 320px;
  overflow-y: scroll;

  border: 1px solid #ccc;
  background-color: #f9f9f9;

  display: flex;
  flex-direction: column;
}

ul::-webkit-scrollbar {
  width: 8px;
}

ul::-webkit-scrollbar-thumb {
  background-color: #ccc;
  cursor: pointer;
}

ul::-webkit-scrollbar-track {
  background-color: #f1f1f1;
}

ul::-webkit-scrollbar-thumb:hover {
  background-color: #999;
}

li:nth-child(2n) {
  background-color: #f0f0f0;
}

li {
  display: grid;
  gap: 8px;
  grid-template-columns: 32px 1fr 80px auto;
  text-align: left;
}

li:hover {
  background-color: #e0e0e0;
}

li .img {
  width: 32px;
  height: 32px;
  display: flex;
  justify-content: center;
  align-items: center;
}

li img {
  width: 16px;
  height: 16px;
}

li span {
  align-self: center;
}

li .link {
  overflow-x: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

li .btn {
  justify-self: flex-end;
  visibility: hidden;
}

li:hover .btn {
  visibility: visible;
}
</style>