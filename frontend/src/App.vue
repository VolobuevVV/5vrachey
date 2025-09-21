<script setup>
import TopPanel from './components/TopPanel.vue'
import HeaderPanel from './components/HeaderPanel.vue'
</script>

<template>
  <div id="app">
    <TopPanel />
    <HeaderPanel />
    <router-view />
  </div>
</template>

<style>
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
}

#app {
  min-height: 100vh;
  min-width: 100vw;
  background-color: rgb(255, 255, 255);
  display: flex;
  flex-direction: column;
  padding-left: env(safe-area-inset-left);
  padding-right: env(safe-area-inset-right);
}
</style>
