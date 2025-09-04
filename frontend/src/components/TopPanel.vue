<template>
  <div :class="{ 'hidden': isScrolled }" class="top-panel"></div>
</template>

<script>
export default {
  name: 'TopPanel',
  data() {
    return {
      isScrolled: false
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    handleScroll() {
      this.isScrolled = window.scrollY > 50
    }
  }
}
</script>

<style scoped>
.top-panel {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 4vh;
  min-height: 30px;
  background-color: rgb(0, 52, 73);
  z-index: 1001;
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.top-panel.hidden {
  opacity: 0;
  transform: translateY(-100%);
  pointer-events: none;
}

</style>
