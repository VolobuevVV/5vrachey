<template>
  <header :class="{ 'scrolled': isScrolled }" class="header-panel">
    <div class="header-content">
      <div class="logo-section">
        <h2 class="logo-text">5 ВРАЧЕЙ</h2>
      </div>

      <div class="right-section">
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

        <nav class="navigation-section">
          <a href="#about">О ЦЕНТРЕ</a>
          <a href="#doctors">Врачи</a>
          <a href="#departments">Отделения</a>
          <a href="#programs">Программы</a>
          <a href="#promotions">Акции</a>
          <a href="#prices">Цены</a>
          <a href="#patients">Пациентам</a>
          <a href="#contacts">Контакты</a>
        </nav>
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

<!-- Стили остаются без изменений -->

<style scoped>
.header-panel {
  position: fixed;
  top: 4vh;
  left: 0;
  width: 100%;
  height: 12vh;
  min-height: 80px;
  background-color: rgb(213, 238, 241);
  transition: all 0.3s ease;
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
}

.header-panel.scrolled {
  top: 0;
  background-color: white !important;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  height: 10vh;
}

.header-content {
  width: 95%;
  max-width: 1400px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo-text {
  margin: 0;
  color: #2c3e50;
  font-weight: 700;
  font-size: 1.8rem;
  white-space: nowrap;
}

.right-section {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 0.5rem;
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
  padding: 0.2rem 0;
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
  border: 2px solid rgb(0, 52, 73);
  border-radius: 25px;
  background: transparent;
  color: rgb(0, 52, 73);
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 0.9rem;
}

.appointment-btn {
  background: rgb(0, 52, 73);
  color: white;
}

.location-btn:hover {
  background: rgba(0, 52, 73, 0.9);
  color: white;
}

.appointment-btn:hover {
  background: rgba(0, 52, 73, 0.9);
  border-color: rgba(0, 52, 73, 0.9);
}

.navigation-section {
  display: flex;
  gap: 1.8rem;
  flex-wrap: wrap;
}

.navigation-section a {
  text-decoration: none;
  color: #2c3e50;
  font-weight: 500;
  font-size: 0.9rem;
  white-space: nowrap;
  transition: text-decoration 0.3s ease;
  padding: 0.2rem 0;
  background: transparent !important;
}

.navigation-section a:hover {
  text-decoration: underline;
}

@media (max-width: 1200px) {
  .header-content {
    flex-direction: column;
    gap: 0.8rem;
  }

  .right-section {
    align-items: center;
    width: 100%;
  }

  .phones-buttons {
    justify-content: center;
  }

  .navigation-section {
    justify-content: center;
    gap: 1.5rem;
  }
}

@media (max-width: 768px) {
  .header-panel {
    top: 4vh;
    height: 14vh;
  }

  .header-panel.scrolled {
    height: 12vh;
  }

  .logo-text {
    font-size: 1.5rem;
    text-align: center;
  }

  .phones-buttons {
    flex-direction: column;
    gap: 0.8rem;
  }

  .phone-numbers {
    align-items: center;
  }

  .phone {
    font-size: 0.8rem;
  }

  .buttons-section {
    gap: 0.6rem;
  }

  .location-btn, .appointment-btn {
    padding: 0.4rem 1rem;
    font-size: 0.8rem;
  }

  .navigation-section {
    gap: 1rem;
    font-size: 0.8rem;
    flex-wrap: nowrap;
    overflow-x: auto;
    padding: 0.5rem 0;
    justify-content: flex-start;
    max-width: 100%;
  }

  .navigation-section a {
    white-space: nowrap;
  }
}
</style>
