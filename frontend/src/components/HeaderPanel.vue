<template>
  <header :class="{ 'scrolled': isScrolled }" class="header-panel">
    <div class="header-container">
      <div class="header-content">
        <div class="logo-section">
          <img
            src="@/assets/svg/logo.svg"
            alt="Медицинский центр 5 ВРАЧЕЙ"
            class="logo-image"
          >
        </div>

        <div class="right-section">
          <div class="top-row">
            <div class="phones-buttons">
              <div class="phone-numbers">
                <div class="phone">+7 (862) 555-14-06</div>
                <div class="phone">+7 (902) 403-55-00</div>
              </div>
              <div class="buttons-section">
                <button class="location-btn" @click="openLocation">Как нас найти</button>
                <button class="appointment-btn" @click="openAppointment">Записаться</button>
              </div>
            </div>
          </div>

          <nav class="navigation-section">
            <a href="#about">О центре</a>
            <a href="#doctors">Врачи</a>
            <a href="#departments">Отделения</a>
            <a href="#prices">Цены</a>
            <a href="#patients">Пациентам</a>
            <a href="#contacts">Контакты</a>
          </nav>
        </div>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'HeaderPanel',
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
    },
    openAppointment() {
      window.open('https://prodoctorov.ru/adler/lpu/47924-pyat-vrachey/', '_blank');
    },
    openLocation() {
      window.open('https://yandex.ru/maps/org/pyat_vrachey/165710051398/?ll=39.927604%2C43.429926&z=16', '_blank', 'width=800,height=600');
    }
  }
}
</script>

<style scoped>
.header-panel {
  position: fixed;
  top: 30px;
  left: 0;
  width: 100%;
  height: 110px;
  background-color: white;
  transition: all 0.3s ease;
  z-index: 1000;
  padding: 0.8rem 0;
  box-shadow: 0 2px 25px rgba(0, 0, 0, 0.08);
}


.header-panel.scrolled {
  top: 0;
  background-color: white !important;
  box-shadow: 0 2px 25px rgba(0, 0, 0, 0.08);
}

.header-container {
  width: 100%;
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 2rem 0;
  display: flex;
  align-items: center;
  height: 100%;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.right-section {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 0.6rem;
  flex: 1;
}

.top-row {
  width: 100%;
  display: flex;
  justify-content: flex-end;
}

.phones-buttons {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

.phone-numbers {
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

.phone {
  font-size: 0.9rem;
  font-weight: 600;
  color: #2c3e50;
  white-space: nowrap;
  transition: all 0.3s ease;
  cursor: pointer;
}

.phone:hover {
  text-decoration: underline;
}

.buttons-section {
  display: flex;
  gap: 0.8rem;
  white-space: nowrap;
}

.location-btn, .appointment-btn {
  padding: 0.5rem 1.2rem;
  border: 2px solid rgb(6, 113, 133);
  border-radius: 20px;
  background: transparent;
  color: rgb(6, 113, 133);
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
}

.appointment-btn {
  background: rgb(6, 113, 133);
  color: white;
}

.location-btn:hover {
  background: rgba(6, 113, 133, 0.9);
  color: white;
}

.appointment-btn:hover {
  background: rgba(6, 113, 133, 0.9);
  border-color: rgba(6, 113, 133, 0.9);
}

.navigation-section {
  display: flex;
  gap: 1.5rem;
  flex-wrap: wrap;
  justify-content: flex-end;
}

.navigation-section a {
  text-decoration: none;
  color: #003449;
  font-weight: 510;
  font-size: 1rem;
  white-space: nowrap;
  transition: all 0.3s ease;
  padding: 0.2rem 0;
  position: relative;
}

.navigation-section a:hover {
  color: rgb(6, 113, 133);
}

.navigation-section a:hover::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 3px;
  background: rgb(6, 113, 133);
}

.logo-section {
  display: flex;
  align-items: center;
}

.logo-image {
  height: 70px;
  width: auto;
  object-fit: contain;
}

</style>
