<template>
  <div class="main-page">
    <div class="content-grid">
      <div class="grid-cell left-cell">
        <Banner />
      </div>

      <div class="grid-cell right-cell">
        <DepartmentsGrid />
      </div>
    </div>
  </div>
</template>

<script>
import Banner from './Banner.vue'
import DepartmentsGrid from './DepartmentsGrid.vue'

export default {
  name: 'MainPage',
  components: {
    Banner,
    DepartmentsGrid
  }
}
</script>

<style scoped>
.main-page {
  padding-top: 160px;
  width: 100%;
  height: 100%;
  background-color: rgb(255, 255, 255);
  display: flex;
  justify-content: center;
  align-items: flex-start;
}

.content-grid {
  display: grid;
  grid-template-columns: 40% 60%;
  width: 100%;
  height: 100%;
}

.grid-cell {
  display: flex;
  justify-content: center;
  align-items: center;
}

</style>
