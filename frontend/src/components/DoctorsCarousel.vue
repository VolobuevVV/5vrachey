<template>
  <section class="doctors-carousel-section" v-if="doctors.length > 0">
    <div class="carousel-container">
      <div class="carousel-header">
        <h2 class="carousel-title">Наши специалисты</h2>
      </div>

      <div class="carousel-content">
        <div class="carousel-controls-left">
          <button class="control-btn prev-btn" @click="prevSlide" :disabled="currentSlide === 0">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M15 18L9 12L15 6" stroke="#003449" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </button>
          <button class="control-btn next-btn" @click="nextSlide" :disabled="currentSlide >= maxSlide">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M9 18L15 12L9 6" stroke="#003449" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            </svg>
          </button>
        </div>

        <div class="carousel-wrapper">
          <div class="doctors-carousel" :style="{ transform: `translateX(-${currentSlide * 100 / slidesToShow}%)` }">
            <DoctorCard
              v-for="doctor in doctors"
              :key="doctor.id"
              :doctor="doctor"
              class="carousel-item"
            />
          </div>
        </div>
      </div>

      <div class="carousel-footer">
        <button class="all-doctors-btn">
          Все врачи
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M5 12H19M19 12L12 5M19 12L12 19" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
          </svg>
        </button>
      </div>
    </div>
  </section>
</template>

<script>
import DoctorCard from './DoctorCard.vue'
import { ref, computed, onMounted } from 'vue'

export default {
  name: 'DoctorsCarousel',
  components: {
    DoctorCard
  },
  setup() {
    const doctors = ref([])
    const currentSlide = ref(0)
    const slidesToShow = 3

    const loadDoctors = async () => {
      try {
        const response = await fetch('/api/doctors')
        doctors.value = await response.json()
      } catch (error) {
        console.error('Ошибка загрузки врачей:', error)
      }
    }

    const maxSlide = computed(() => {
      return Math.max(0, Math.ceil(doctors.value.length / slidesToShow) - 1)
    })

    const nextSlide = () => {
      if (currentSlide.value < maxSlide.value) {
        currentSlide.value++
      }
    }

    const prevSlide = () => {
      if (currentSlide.value > 0) {
        currentSlide.value--
      }
    }

    onMounted(() => {
      loadDoctors()
    })

    return {
      doctors,
      currentSlide,
      maxSlide,
      slidesToShow,
      nextSlide,
      prevSlide
    }
  }
}
</script>

<style scoped>
.doctors-carousel-section {
  padding: 4rem 0;
  background: white;
}

.carousel-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.carousel-header {
  margin-bottom: 3rem;
  text-align: center;
}

.carousel-title {
  color: rgb(0, 52, 73);
  font-size: 2.2rem;
  font-weight: 700;
  margin: 0;
}

.carousel-content {
  display: flex;
  align-items: flex-start;
  gap: 2rem;
  margin-bottom: 3rem;
}

.carousel-controls-left {
  display: flex;
  gap: 1rem;
  padding-top: 2rem;
}

.control-btn {
  width: 48px;
  height: 48px;
  border: 2px solid rgb(0, 52, 73);
  border-radius: 50%;
  background: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
}

.control-btn:hover:not(:disabled) {
  background: rgb(0, 52, 73);
}

.control-btn:hover:not(:disabled) svg path {
  stroke: white;
}

.control-btn:disabled {
  border-color: #ccc;
  cursor: not-allowed;
}

.control-btn:disabled svg path {
  stroke: #ccc;
}

.carousel-wrapper {
  overflow: hidden;
  flex: 1;
  padding: 1rem 0;
}

.doctors-carousel {
  display: flex;
  transition: transform 0.5s ease;
  gap: 2rem;
}

.carousel-item {
  flex: 0 0 calc(33.333% - 1.33rem);
  min-width: 0;
}

.carousel-item :deep(.doctor-card) {
  border: 2px solid rgb(6, 113, 133);
  border-radius: 15px;
}

.carousel-footer {
  display: flex;
  justify-content: center;
}

.all-doctors-btn {
  font-size: 1rem;
  font-weight: 700;
  color: rgb(0, 52, 73);
  padding: 0.8rem 2.5rem;
  border: 2px solid rgb(0, 52, 73);
  border-radius: 50px;
  transition: all 0.3s ease;
  cursor: pointer;
  white-space: nowrap;
  background: white;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.all-doctors-btn:hover {
  background: rgb(0, 52, 73);
  color: white;
}
</style>
