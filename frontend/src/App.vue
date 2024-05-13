<script lang="ts" setup>
import { onMounted } from 'vue'
import { InitDatabase } from '../wailsjs/go/main/App'

onMounted(async () => {
  await InitDatabase()
})
</script>

<template>
  <router-view />
</template>

<style>
#app {
  display: flex;
}

p {
  margin: 0;
}
</style>