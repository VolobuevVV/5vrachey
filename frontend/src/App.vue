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
/* Скрыть скроллбары но оставить функциональность скролла */
::-webkit-scrollbar {
  display: none;
}

html {
  scrollbar-width: none;
}

body {
  -ms-overflow-style: none;
}

html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  background-color: rgb(255, 255, 255);
}

#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  width: 100%;
  margin: 0;
  padding: 0;
}

/* Глобальный box-sizing */
*, *::before, *::after {
  box-sizing: border-box;
}
</style>
