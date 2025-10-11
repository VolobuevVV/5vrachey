<template>
  <section class="hero-section">
    <div class="text-container">
      <div class="text-side">
        <h1 class="hero-title">Диагностика</h1>
        <p class="hero-subtitle">Мировые стандарты диагностики. Полный спектр услуг от простых до самых сложных</p>
        <div class="hero-buttons">
          <button class="hero-appointment-btn" @click="openAppointment">ЗАПИСАТЬСЯ</button>
          <div class="hero-phone" @click="openAppointment">+7 (862) 555-14-06</div>
        </div>
      </div>
    </div>
    <div class="background-overlay"></div>
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
  width: 90%;
  margin-left: 10%;
  height: 100%;
  background-image: url('@/assets/images/banner3.jpg');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  position: relative;
  border-radius: 20px;
  display: flex;
  align-items: flex-end;
  justify-content: center;
  overflow: hidden;
}

.text-container {
  position: relative;
  width: 100%;
  max-width: 800px;
  display: flex;
  justify-content: center;
}

.text-side {
  position: relative;
  z-index: 2;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  padding: 0 20px 2rem 20px;
  margin-bottom: 0;
  width: 100%;
}

.background-overlay {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 60%; /* уменьшил высоту */
  background: radial-gradient(
    ellipse 70% 60% at center 60%, /* сплюснул и сместил вниз */
    rgba(255, 255, 255, 0.8) 0%,
    rgba(255, 255, 255, 0.8) 20%,
    rgba(255, 255, 255, 0.8) 40%,
    rgba(255, 255, 255, 0.7) 50%,
    rgba(255, 255, 255, 0.6) 60%,
    rgba(255, 255, 255, 0.3) 80%,
    rgba(255, 255, 255, 0) 100%
  );
  z-index: 1;
  pointer-events: none;
}

.hero-title {
  font-size: 2.4rem;
  font-weight: 800;
  color: rgb(0, 52, 73);
  margin-bottom: 0.5rem;
  line-height: 1;
}

.hero-subtitle {
  font-size: 1rem;
  color: #2c3e50;
  margin-bottom: 1rem;
  line-height: 1.3;
  font-weight: 400;
  width: 100%;
}

.hero-buttons {
  display: flex;
  gap: 0.6rem;
  align-items: center;
  flex-wrap: wrap;
  justify-content: center;
}

.hero-appointment-btn {
  padding: 0.7rem 1.4rem;
  background: rgb(6, 113, 133);
  color: white;
  border: none;
  border-radius: 50px;
  font-size: 0.85rem;
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
  font-size: 0.85rem;
  font-weight: 700;
  color: rgb(0, 52, 73);
  padding: 0.5rem 1.1rem;
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
