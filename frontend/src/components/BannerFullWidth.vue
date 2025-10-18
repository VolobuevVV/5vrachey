<template>
  <section class="fullwidth-banner" :style="{ backgroundImage: `url(${bannerImage})` }">
    <div class="banner-content">
      <div class="text-content">
        <h1 class="banner-title">Диагностика</h1>
        <p class="banner-subtitle">Мировые стандарты диагностики. Полный спектр услуг от простых до самых сложных</p>
        <div class="banner-buttons">
          <button class="appointment-btn" @click="openAppointment">ЗАПИСАТЬСЯ</button>
          <div class="phone-number" @click="openAppointment">+7 (862) 555-14-06</div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'BannerFullWidth',
  data() {
    const isLocal = window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1'
    return {
      bannerImage: isLocal ? 'http://localhost:8081/banners/banner2.webp' : '/files/banners/banner2.webp'
    }
  },
  methods: {
    openAppointment() {
      window.open('https://prodoctorov.ru/adler/lpu/47924-pyat-vrachey/', '_blank');
    }
  }
}
</script>

<style scoped>
.fullwidth-banner {
  width: 100%;
  height: 400px;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.banner-content {
  position: relative;
  width: 100%;
  max-width: 1200px;
  display: flex;
  justify-content: flex-start;
  padding-left: 20px;
}

.text-content {
  position: relative;
  z-index: 2;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  width: 100%;
  max-width: 600px;
}

.banner-title {
  font-size: 3rem;
  font-weight: 800;
  color: rgb(0, 52, 73);
  margin-bottom: 1rem;
  line-height: 1.1;
}

.banner-subtitle {
  font-size: 1.2rem;
  color: #2c3e50;
  margin-bottom: 2rem;
  line-height: 1.4;
  font-weight: 400;
  width: 100%;
}

.banner-buttons {
  display: flex;
  gap: 1rem;
  align-items: center;
  flex-wrap: wrap;
  justify-content: center;
}

.appointment-btn {
  padding: 1rem 2rem;
  background: rgb(6, 113, 133);
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

.appointment-btn:hover {
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

.phone-number {
  font-size: 1rem;
  font-weight: 700;
  color: rgb(0, 52, 73);
  padding: 0.8rem 1.5rem;
  border: 2px solid rgb(0, 52, 73);
  border-radius: 50px;
  transition: all 0.3s ease;
  cursor: pointer;
  white-space: nowrap;
}

.phone-number:hover {
  background: rgb(0, 52, 73);
  color: white;

}
</style>
