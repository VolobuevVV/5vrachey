<template>
  <section class="hero-section">
    <div class="hero-main">
      <div class="hero-content">
        <div class="text-side">
          <h1 class="hero-title">Диагностика</h1>
          <p class="hero-subtitle">Мировые стандарты диагностики. Полный спектр услуг от простых до самых сложных</p>
          <div class="hero-buttons">
            <button class="hero-appointment-btn" @click="openAppointment">ЗАПИСАТЬСЯ</button>
            <div class="hero-phone" @click="openAppointment">+7 (862) 555-14-06</div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'Banner',
  methods: {
    openAppointment() {
      window.open('https://prodoctorov.ru/adler/lpu/47924-pyat-vrachey/', '_blank');
    }
  }
}
</script>

<style scoped>
.hero-section {
  width: 100%;
  background: transparent;
  position: relative;
}

.hero-main {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0;
}

.hero-content {
  display: grid;
  grid-template-columns: 1fr;
  align-items: center;
  width: 100%;
}

.text-side {
  padding-right: 0;
}

.hero-title {
  font-size: 2.8rem;
  font-weight: 800;
  color: rgb(0, 52, 73);
  margin-bottom: 1.2rem;
  line-height: 1.1;
}

.hero-subtitle {
  font-size: 1.1rem;
  color: #2c3e50;
  margin-bottom: 2rem;
  line-height: 1.5;
  font-weight: 400;
  width: 100%;
}

.hero-buttons {
  display: flex;
  gap: 1rem;
  align-items: center;
  flex-wrap: wrap;
}

.hero-appointment-btn {
  padding: 1.1rem 2.2rem;
  background: rgb(6, 113, 130);
  color: white;
  border: none;
  border-radius: 50px;
  font-size: 1rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s ease;
  text-transform: uppercase;
  white-space: nowrap;
  animation: pulse 2s infinite;
}

.hero-appointment-btn:hover {
  background: rgba(0, 52, 73, 0.9);
  transform: translateY(-2px);
  animation: none;
}

@keyframes pulse {
  0% {
    transform: scale(1);
    box-shadow: 0 0 0 0 rgba(6, 113, 133, 0.7);
  }
  50% {
    transform: scale(1.03);
    box-shadow: 0 0 0 10px rgba(6, 113, 133, 0);
  }
  100% {
    transform: scale(1);
    box-shadow: 0 0 0 0 rgba(6, 113, 133, 0);
  }
}

.hero-phone {
  font-size: 1rem;
  font-weight: 700;
  color: rgb(0, 52, 73);
  padding: 0.8rem 1.6rem;
  border: 2px solid rgb(0, 52, 73);
  border-radius: 50px;
  transition: all 0.3s ease;
  cursor: pointer;
  white-space: nowrap;
}

.hero-phone:hover {
  background: rgb(0, 52, 73);
  color: white;
}
</style>
