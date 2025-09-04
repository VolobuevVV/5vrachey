<template>
  <section class="hero-section">

    <!-- Основной контент HeroSection -->
    <div class="hero-main">
      <div class="hero-container">
        <div class="hero-content">
          <div class="text-side">
            <h1 class="hero-title">Диагностика</h1>
            <p class="hero-subtitle">Мировые стандарты диагностики. Полный спектр услуг от простых до самых сложных</p>
            <div class="hero-buttons">
              <button class="hero-appointment-btn" @click="openAppointment">ЗАПИСАТЬСЯ</button>
              <div class="hero-phone" @click="openAppointment">+7 (862) 555-14-06</div>
            </div>
          </div>
          <div class="image-side">
            <div class="white-square">
              <div class="placeholder">Изображение</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Виджеты под основным контентом -->
    <div class="hero-widgets">
      <!-- Виджет 1: О центре -->
      <div class="hero-widget">
        <div class="widget-content">
          <h2>О медицинском центре "5 ВРАЧЕЙ"</h2>
          <p>Современный медицинский центр, оснащенный передовым оборудованием и укомплектованный высококвалифицированными специалистами</p>
        </div>
      </div>

      <!-- Виджет 2: Услуги -->
      <div class="hero-widget">
        <div class="widget-content">
          <h2>Наши услуги</h2>
          <div class="services-grid">
            <div class="service-item">✅ Диагностика</div>
            <div class="service-item">✅ Консультации</div>
            <div class="service-item">✅ Анализы</div>
            <div class="service-item">✅ Лечение</div>
            <div class="service-item">✅ Реабилитация</div>
            <div class="service-item">✅ Профилактика</div>
          </div>
        </div>
      </div>

      <!-- Виджет 3: Врачи -->
      <div class="hero-widget">
        <div class="widget-content">
          <h2>Наши специалисты</h2>
          <p>В нашем центре работают врачи высшей категории с многолетним опытом работы</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: 'HeroSection',
  methods: {
    openAppointment() {
      window.open('https://prodoctorov.ru/adler/lpu/47924-pyat-vrachey/', '_blank');
    }
  }
}
</script>

<style scoped>
/* Основной контент HeroSection */
.hero-section {
  width: 100%;
  background: rgb(213, 238, 241);
  min-height: 100vh;
  position: relative;
}

.hero-main {
  min-height: 80vh; /* Остальные 80% */
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  margin: 0;
}

.hero-container {
  width: 100%;
  margin: 0 auto;
  padding: 0 5vw;
}

.hero-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
  align-items: center;
  width: 100%;
}

.text-side {
  padding-right: 1rem;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 800;
  color: rgb(0, 52, 73);
  margin-bottom: 1.5rem;
  line-height: 1.1;
}

.hero-subtitle {
  font-size: 1.3rem;
  color: #2c3e50;
  margin-bottom: 2.5rem;
  line-height: 1.5;
  font-weight: 400;
  width: 100%;
}

.hero-buttons {
  display: flex;
  gap: 1.2rem;
  align-items: center;
  flex-wrap: nowrap;
}

.hero-appointment-btn {
  padding: 1.1rem 2.2rem;
  background: rgb(0, 52, 73);
  color: white;
  border: none;
  border-radius: 50px;
  font-size: 1rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.3s ease;
  text-transform: uppercase;
  white-space: nowrap;
}

.hero-appointment-btn:hover {
  background: rgba(0, 52, 73, 0.9);
  transform: translateY(-2px);
}

.hero-phone {
  font-size: 1.1rem;
  font-weight: 700;
  color: rgb(0, 52, 73);
  padding: 0.9rem 1.8rem;
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

.image-side {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.white-square {
  width: 50vw;
  height: 64vh;
  background: white;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 3px rgba(0, 0, 0, 0.04);
  margin-right: -1rem;
}

.placeholder {
  color: #ccc;
  font-size: 1.2rem;
  font-weight: 500;
}

/* Стили для виджетов */
.hero-widgets {
  width: 100%;
  padding: 0;
}

.hero-widget:first-child {
  margin-top: -10vh; /* Поднимает только первый виджет */
  padding-top: 15vh; /* Уменьшает отступ сверху */
}

.hero-widget {
  padding: 4rem 0;
  width: 100%;
  background: rgb(213, 238, 241);
}

.hero-widget:nth-child(even) {
  background: rgba(255, 255, 255, 0.5);
}

.widget-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
}

.widget-content h2 {
  color: rgb(0, 52, 73);
  font-size: 2.5rem;
  margin-bottom: 1.5rem;
  text-align: center;
}

.widget-content p {
  color: #2c3e50;
  font-size: 1.2rem;
  line-height: 1.6;
  text-align: center;
  max-width: 800px;
  margin: 0 auto;
}

.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
  margin-top: 2rem;
}

.service-item {
  background: white;
  padding: 1.5rem;
  border-radius: 10px;
  text-align: center;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
  color: rgb(0, 52, 73);
  font-weight: 600;
  font-size: 1.1rem;
}


@media (max-width: 1024px) {

  .hero-content {
    grid-template-columns: 1fr;
    gap: 2.5rem;
    text-align: center;
  }

  .text-side {
    padding-right: 0;
  }

  .image-side {
    justify-content: center;
  }

  .hero-title {
    font-size: 2.8rem;
  }

  .hero-buttons {
    justify-content: center;
  }
}

@media (max-width: 768px) {

  .hero-main {
    min-height: 110vh;
  }

  .hero-container {
    padding: 0 1rem;
  }

  .hero-title {
    font-size: 2.3rem;
  }

  .hero-widget {
    padding: 2rem 0;
  }

  .widget-content {
    padding: 0 1rem;
  }

  .widget-content h2 {
    font-size: 2rem;
  }

  .widget-content p {
    font-size: 1.1rem;
  }

  .services-grid {
    grid-template-columns: 1fr;
  }

  .service-item {
    padding: 1.2rem;
  }
}
</style>
