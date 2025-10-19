<template>
  <div class="main-page">
    <BannerFullWidth/>
    <div class="content-section">
      <Tabs />
      <AboutSection />
    </div>
    <FooterPanel />
  </div>
</template>

<script>
import Banner from './Banner.vue'
import DepartmentsGrid from './DepartmentsGrid.vue'
import FooterPanel from './FooterPanel.vue'
import BannerFullWidth from "@/components/BannerFullWidth.vue";
import Tabs from "@/components/Tabs.vue";
import AboutSection from "@/components/AboutSection.vue";

export default {
  name: 'MainPage',
  components: {
    AboutSection,
    Tabs,
    BannerFullWidth,
    Banner,
    DepartmentsGrid,
    FooterPanel
  }
}
</script>

<style scoped>
.main-page {
  padding-top: 140px;
  background-color: rgb(255, 255, 255);
  height: 100vh;
  width: 100vw;
  margin: 0;
  padding-left: 0;
  padding-right: 0;
}


.content-grid {
  display: grid;
  grid-template-columns: 40% 60%;
  height: 380px;
  margin-bottom: 500px;
  width: 100%;
}

.grid-cell {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}
</style>
