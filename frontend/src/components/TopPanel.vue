<template>
  <div :class="{ 'hidden': isScrolled }" class="top-panel">
    <div class="discount-text">Скидка 20% на первое посещение</div>
  </div>
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
  min-height: 20px;
  background-color: rgb(6, 113, 133);
  z-index: 1001;
  transition: opacity 0.3s ease, transform 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.discount-text {
  color: white;
  font-size: 0.8rem;
  font-weight: 680;
  text-align: center;
  white-space: nowrap;
}

.top-panel.hidden {
  opacity: 0;
  transform: translateY(-100%);
  pointer-events: none;
}

/* Адаптивность */
@media (max-width: 768px) {
  .discount-text {
    font-size: 0.7rem;
  }
}

@media (max-width: 480px) {
  .discount-text {
    font-size: 0.7rem;
    white-space: normal;
    padding: 0 0.5rem;
    line-height: 1.2;
  }

  .top-panel {
    min-height: 25px;
    height: 4vh;
  }
}
</style>
