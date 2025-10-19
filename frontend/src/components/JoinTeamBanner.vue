<template>
  <section class="join-team-banner" :style="{ backgroundImage: `url(${bannerImage})` }">
    <div class="banner-content">
      <div class="text-content">
        <h2 class="banner-title">Присоединяйтесь к команде "Пять Врачей"</h2>
        <p class="banner-subtitle">Станьте частью современной медицинской команды, где ценятся профессионализм, развитие и забота о пациентах</p>
        <div class="banner-buttons">
          <button class="join-btn" @click="openVacancies">ВАКАНСИИ</button>
          <button class="contact-btn" @click="openContact">СВЯЗАТЬСЯ С НАМИ</button>
        </div>
      </div>
    </div>
    <div class="banner-overlay"></div>
  </section>
</template>

<script>
export default {
  name: 'JoinTeamBanner',
  data() {
    const isLocal = window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1'
    return {
      bannerImage: isLocal ? 'http://localhost:8081/banners/team.webp' : '/files/banners/team.webp'
    }
  },
  methods: {
    openVacancies() {
      // Открыть страницу с вакансиями
      console.log('Открыть вакансии')
    },
    openContact() {
      // Открыть контакты
      console.log('Открыть контакты')
    }
  }
}
</script>

<style scoped>
.join-team-banner {
  width: 100%;
  height: 400px;
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  overflow: hidden;
}

.banner-content {
  position: relative;
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: flex-start;
  padding-left: 2rem;
  z-index: 2;
}

.text-content {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-align: left;
  max-width: 600px;
}

.banner-title {
  font-size: 2.5rem;
  font-weight: 800;
  color: rgb(0, 52, 73);
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

.banner-subtitle {
  font-size: 1.2rem;
  color: #2c3e50;
  margin-bottom: 2.5rem;
  line-height: 1.5;
  font-weight: 400;
  width: 100%;
}

.banner-buttons {
  display: flex;
  gap: 1.5rem;
  align-items: center;
  flex-wrap: wrap;
}

.join-btn {
  padding: 1rem 2.5rem;
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

.join-btn:hover {
  background: rgba(0, 52, 73, 0.9);
  transform: translateY(-2px);
  animation: none;
}

.contact-btn {
  padding: 1rem 2rem;
  background: white;
  color: rgb(0, 52, 73);
  border: 2px solid rgb(0, 52, 73);
  border-radius: 50px;
  font-size: 1rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s ease;
  text-transform: uppercase;
  white-space: nowrap;
}

.contact-btn:hover {
  background: rgb(0, 52, 73);
  color: white;
  transform: translateY(-2px);
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

.banner-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    90deg,
    rgba(255, 255, 255, 0.9) 0%,
    rgba(255, 255, 255, 0.7) 30%,
    rgba(255, 255, 255, 0.4) 70%,
    rgba(255, 255, 255, 0) 100%
  );
  z-index: 1;
  pointer-events: none;
}
</style>
