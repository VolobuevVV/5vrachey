<template>
  <div :class="{ 'hidden': isScrolled }" class="top-panel">
    <div class="discount-text">{{ discountText }}</div>
  </div>
</template>

<script>
export default {
  name: 'TopPanel',
  data() {
    return {
      isScrolled: false,
      discountText: 'Загрузка...'
    }
  },
  async mounted() {
    window.addEventListener('scroll', this.handleScroll)
    await this.fetchDiscountText()
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    handleScroll() {
      this.isScrolled = window.scrollY > 50
    },
    async fetchDiscountText() {
      try {
        const response = await fetch('/api/top-panel-text')
        const data = await response.json()
        this.discountText = data.text
      } catch (error) {
        console.error('Ошибка загрузки текста:', error)
        this.discountText = 'Скидка 20% на первое посещение'
      }
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
  height: 30px;
  background-color: rgb(6, 113, 133);
  z-index: 1001;
  transition: opacity 0.3s ease, transform 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.discount-text {
  color: white;
  font-size: 14px;
  font-weight: 680;
  text-align: center;
  white-space: nowrap;
}

.top-panel.hidden {
  opacity: 0;
  transform: translateY(-100%);
  pointer-events: none;
}
</style>
