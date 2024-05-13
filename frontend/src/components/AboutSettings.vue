<script setup lang="ts">
import { Icon } from '@iconify/vue';
</script>

<template>
  <div>
    <h2>MrRss</h2>
    <h4>Versatile, Light-weight, AI-powered RSS Reader</h4>
  </div>
  <img src="../assets/images/appicon.png" class="logo" alt="MrRss Logo" />
  <div class="version">
    <button class="btn" title="Check update">
      <Icon icon="material-symbols:deployed-code-update-outline" />
    </button>
    <p>
      Current version: 0.1.0
    </p>
  </div>
  <div class="copyright">
    <p>© 2024 <a href="https://github.com/WCY-dt/MrRSS" rel="help" target="_blank">MrRss</a>. All rights reserved.</p>
    <p>Designed and developed by <a href="https://ch3nyang.top/about/" rel="help" target="_blank">Ch3nyang</a>.</p>
  </div>
</template>

<style scoped>
h2 {
  margin: 8px 0;
}

h4 {
  margin: 0;
  font-weight: 400;
  color: #666;
}

.logo {
  min-width: 240px;
  max-width: 240px;
  height: 240px;
  margin: 0 auto;
  display: block;
}

.version {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.btn {
  width: 32px;
  height: 32px;
  font-size: large;
  color: black;
  background-color: #e0e0e0;
  border: none;
  border-radius: 50%;
  cursor: pointer;

  display: flex;
  justify-content: center;
  align-items: center;
}

.btn:hover {
  background-color: #d0d0d0;
}

.copyright {
  display: flex;
  flex-direction: column;
  align-items: center;

  font-size: small;
}

a {
  color: #000;
}
</style>