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
  padding-top: 21vh;
  width: 100vw;
  min-height: 100vh;
  background-color: rgb(255, 255, 255);
  display: flex;
  justify-content: center;
  align-items: flex-start;
}

.content-grid {
  display: grid;
  grid-template-columns: 35% 65%;
  gap: 2rem;
  width: 90vw;
  max-width: 1400px;
  margin: 2vh 0 0;
  padding: 0;
  height: 400px;
}

.grid-cell {
  display: flex;
  align-items: center;
  height: 400px;
}

.left-cell {
  padding: 0;
  text-align: left;
  justify-self: start;
}

.right-cell {
  flex-direction: column;
  justify-content: center;
  padding: 2rem;
}
</style>
