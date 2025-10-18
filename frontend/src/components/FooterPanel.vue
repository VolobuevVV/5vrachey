<template>
  <footer class="footer-panel">
    <div class="footer-content">

      <!-- Колонка 1: Контакты -->
      <div class="footer-column">
        <h4>Контакты</h4>
        <div class="contact-item">
          <img src="@/assets/icons/phone.svg" alt="Телефон">
          <div class="phones">
            <div>+7 (862) 555-14-06</div>
            <div>+7 (902) 403-55-00</div>
          </div>
        </div>
        <div class="contact-item">
          <img src="@/assets/icons/location.svg" alt="Адрес">
          <div>г. Адлер, ул. Примерная, 123</div>
        </div>
        <div class="contact-item">
          <img src="@/assets/icons/clock.svg" alt="Время работы">
          <div>Пн-Пт: 8:00-20:00<br>Сб-Вс: 9:00-18:00</div>
        </div>
      </div>

      <!-- Колонка 2: Отделения -->
      <div class="footer-column">
        <h4>Отделения</h4>
        <a href="#therapy">Терапия</a>
        <a href="#surgery">Хирургия</a>
        <a href="#diagnostics">Диагностика</a>
        <a href="#pediatrics">Педиатрия</a>
      </div>

      <!-- Колонка 3: Документы -->
      <div class="footer-column">
        <h4>Информация</h4>
        <a href="#licenses">Лицензии</a>
        <a href="#docs">Документы</a>
        <a href="#privacy">Политика конфиденциальности</a>
        <button class="online-appointment" @click="openAppointment">
          Записаться онлайн
        </button>
      </div>

      <!-- Колонка 4: Соцсети -->
      <div class="footer-column">
        <h4>Мы в соцсетях</h4>
        <div class="social-links">
          <a href="#"><img src="@/assets/icons/vk.svg" alt="VK"></a>
          <a href="#"><img src="@/assets/icons/telegram.svg" alt="Telegram"></a>
          <a href="#"><img src="@/assets/icons/whatsapp.svg" alt="WhatsApp"></a>
        </div>
        <div class="footer-badge">
          <img src="@/assets/icons/medal.svg" alt="Рейтинг">
          <span>Топ-3 клиник Адлера</span>
        </div>
      </div>

    </div>

    <!-- Нижняя строка -->
    <div class="footer-bottom">
      <div class="copyright">© 2024 Медицинский центр "5 ВРАЧЕЙ". Все права защищены.</div>
      <div class="developers">Разработано с ♥ для вашего здоровья</div>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'FooterPanel',
  methods: {
    openAppointment() {
      window.open('https://prodoctorov.ru/adler/lpu/47924-pyat-vrachey/', '_blank');
    }
  }
}
</script>

<style scoped>
.footer-panel {
  background: rgb(255, 255, 255);
  color: #003449;
  padding: 3rem 2rem 1rem;
  margin-top: auto;
  width: 100%;
}

.footer-content {
  max-width: 1400px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 2rem;
}

.footer-column h4 {
  color: rgb(6, 113, 133);
  margin-bottom: 1rem;
  font-size: 1.2rem;
  font-weight: 700;
}

.contact-item {
  display: flex;
  align-items: flex-start;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.contact-item img {
  width: 20px;
  height: 20px;
  margin-top: 0.2rem;
  color: rgb(6, 113, 133);
}

.footer-column a {
  display: block;
  color: #2c3e50;
  text-decoration: none;
  margin-bottom: 0.5rem;
  transition: color 0.3s;
  font-weight: 510;
}

.footer-column a:hover {
  color: rgb(6, 113, 133);
}

.online-appointment {
  background: rgb(6, 113, 133);
  color: white;
  border: none;
  padding: 0.7rem 1.5rem;
  border-radius: 25px;
  font-weight: 600;
  cursor: pointer;
  margin-top: 1rem;
  transition: all 0.3s ease;
  font-size: 0.9rem;
  border: 2px solid rgb(6, 113, 133);
}

.online-appointment:hover {
  background: rgba(6, 113, 133, 0.9);
  border-color: rgba(6, 113, 133, 0.9);
  transform: translateY(-1px);
}

.social-links {
  display: flex;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.social-links img {
  width: 32px;
  height: 32px;
  transition: transform 0.3s;
  color: rgb(6, 113, 133);
}

.social-links img:hover {
  transform: scale(1.1);
}

.footer-badge {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: rgba(6, 113, 133, 0.1);
  padding: 0.5rem 1rem;
  border-radius: 10px;
  color: #003449;
  font-weight: 600;
}

.footer-badge img {
  width: 24px;
  height: 24px;
  color: rgb(6, 113, 133);
}

.footer-bottom {
  max-width: 1400px;
  margin: 2rem auto 0;
  padding-top: 1rem;
  border-top: 1px solid rgba(0, 52, 73, 0.1);
  display: flex;
  justify-content: space-between;
  color: #666;
  font-size: 0.9rem;
}
</style>
