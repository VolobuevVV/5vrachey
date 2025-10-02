<template>
  <header :class="{ 'scrolled': isScrolled, 'menu-open': isMenuOpen }" class="header-panel">
    <div class="header-container">
      <div class="header-content">
        <div class="logo-section">
          <img
            src="@/assets/svg/logo.svg"
            alt="Медицинский центр 5 ВРАЧЕЙ"
            class="logo-image"
          >
        </div>

        <button class="burger-menu" @click="toggleMenu">
          <span></span>
          <span></span>
          <span></span>
        </button>

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
            <a href="#about" @click="closeMenu">О центре</a>
            <a href="#doctors" @click="closeMenu">Врачи</a>
            <a href="#departments" @click="closeMenu">Отделения</a>
            <a href="#prices" @click="closeMenu">Цены</a>
            <a href="#patients" @click="closeMenu">Пациентам</a>
            <a href="#contacts" @click="closeMenu">Контакты</a>
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
      isScrolled: false,
      isMenuOpen: false
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
    window.addEventListener('resize', this.handleResize)
    this.handleResize()
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll)
    window.removeEventListener('resize', this.handleResize)
  },
  methods: {
    handleScroll() {
      this.isScrolled = window.scrollY > 50
    },
    handleResize() {
      if (window.innerWidth > 768) {
        this.isMenuOpen = false
      }
    },
    toggleMenu() {
      this.isMenuOpen = !this.isMenuOpen
    },
    closeMenu() {
      this.isMenuOpen = false
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
  top: 4vh;
  left: 0;
  width: 100%;
  height: 17vh;
  min-height: 100px;
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
  padding: 0 2rem;
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.burger-menu {
  display: none;
  flex-direction: column;
  gap: 3px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.4rem;
  z-index: 1002;
}

.burger-menu span {
  width: 22px;
  height: 3px;
  background: rgb(6, 113, 133);
  transition: all 0.3s ease;
  transform-origin: center;
}

.header-panel.menu-open .burger-menu span:nth-child(1) {
  transform: rotate(45deg) translate(5px, 5px);
}

.header-panel.menu-open .burger-menu span:nth-child(2) {
  opacity: 0;
}

.header-panel.menu-open .burger-menu span:nth-child(3) {
  transform: rotate(-45deg) translate(5px, -5px);
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
  height: 10vh;
  width: auto;
  object-fit: contain;
  transition: all 0.3s ease;
}

.header-panel.scrolled .logo-image {
  height: 8vh;
}

@media (max-width: 768px) {
  .header-panel {
    top: 3vh;
    padding: 0.6rem 0;
    min-height: 60px;
  }

  .header-panel.scrolled {
    padding: 0.3rem 0;
  }

  .header-container {
    padding: 0 1rem;
  }

  .burger-menu {
    display: flex;
  }

  .right-section {
    position: fixed;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100vh;
    background: white;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 1.5rem;
    transition: left 0.3s ease;
    z-index: 1001;
    padding: 1.5rem;
  }

  .header-panel.menu-open .right-section {
    left: 0;
  }

  .top-row {
    justify-content: center;
  }

  .phones-buttons {
    flex-direction: column;
    gap: 1.2rem;
    align-items: center;
  }

  .phone-numbers {
    align-items: center;
    text-align: center;
    gap: 0.3rem;
  }

  .phone {
    font-size: 1rem;
  }

  .buttons-section {
    flex-direction: column;
    gap: 0.8rem;
  }

  .location-btn, .appointment-btn {
    padding: 0.8rem 1.5rem;
    font-size: 1rem;
    width: 220px;
    text-align: center;
  }

  .navigation-section {
    flex-direction: column;
    gap: 1.2rem;
    text-align: center;
  }

  .navigation-section a {
    font-size: 2rem;
    padding: 0.4rem 0;
  }

  .logo-image {
    height: 10vh;
  }

  .header-panel.scrolled .logo-image {
    height: 8vh;
  }
}
</style>
