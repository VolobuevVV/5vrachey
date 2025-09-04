<script setup>
import TopPanel from './components/TopPanel.vue'
import HeaderPanel from './components/HeaderPanel.vue'
import HeroSection from './components/HeroSection.vue'
</script>

<template>
  <TopPanel />
  <HeaderPanel />

  <main class="main-content">
    <HeroSection />
  </main>
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
  width: 100%;
  display: flex;
  flex-direction: column;
  background-color: rgb(213, 238, 241);
}

.main-content {
  min-height: 100vh;
}

/* Адаптивность */
@media (max-width: 1920px) {
  .main-content {
    padding-top: 100vh;
  }
}

@media (max-width: 1400px) {
  .main-content {
    padding-top: 175vh;
  }
}

@media (max-width: 768px) {
  .main-content {
    padding-top: 0vh;
  }
}
</style>
