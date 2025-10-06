<template>
  <div>
    <div class="doctor-card">
      <div class="doctor-image">
        <img :src="doctorPhoto" :alt="fullName" class="doctor-photo">
      </div>
      <div class="doctor-info">
        <h3 class="doctor-name">{{ fullName }}</h3>
        <p class="doctor-position">{{ formattedPositions }}</p>
        <p class="doctor-experience">Опыт работы: {{ doctor.experience }} лет</p>

        <!-- Если запись активна - кнопка записи -->
        <button
          class="appointment-btn"
          @click="openAppointment"
          v-if="doctor.available_for_appointment"
        >
          Записаться онлайн
        </button>

        <!-- Если запись неактивна - виджет с телефоном -->
        <div v-else class="phone-widget">
          <div class="phone-icon">📞</div>
          <div class="phone-info">
            <div class="phone-text">Для записи позвоните:</div>
            <div class="phone-number">+7 (862) 555-14-06</div>
            <div class="phone-number">+7 (902) 403-55-00</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Модальное окно с iframe -->
    <div v-if="showIframe" class="iframe-modal">
      <div class="modal-content">
        <div class="modal-header">
          <h3>Запись к {{ fullName }}</h3>
          <button class="close-btn" @click="showIframe = false">×</button>
        </div>
        <iframe
          :src="appointmentLink"
          class="appointment-iframe"
          frameborder="0"
          allowfullscreen
        ></iframe>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DoctorCard',
  props: {
    doctor: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      showIframe: false
    }
  },
  computed: {
    fullName() {
      return `${this.doctor.last_name} ${this.doctor.first_name} ${this.doctor.patronymic}`
    },
    formattedPositions() {
      return this.doctor.positions.join(', ')
    },
    doctorPhoto() {
      return `/files/${this.doctor.id}.webp`
    },
    appointmentLink() {
      // Извлекаем ID до '-' (например из "629306-vasileva" берем "629306")
      const doctorId = this.doctor.id.split('-')[0]
      return `https://booking.medflex.ru?user=e6b1637666264d8202e6253e62b38aeb&employeeId=${doctorId}&source=4`
    }
  },
  methods: {
    openAppointment() {
      this.showIframe = true
    }
  }
}
</script>

<style scoped>
/* Стили для виджета телефона */
.phone-widget {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 10px;
  border: 1px solid #e9ecef;
}

.phone-icon {
  font-size: 1.5rem;
}

.phone-info {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}

.phone-text {
  font-size: 0.9rem;
  color: #6c757d;
  font-weight: 500;
}

.phone-number {
  font-size: 1rem;
  color: #003449;
  font-weight: 600;
  text-decoration: none;
}

.phone-number:hover {
  text-decoration: underline;
}

/* Остальные стили без изменений */
.doctor-card {
  background: white;
  border-radius: 15px;
  padding: 1.5rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  display: flex;
  gap: 1.5rem;
  transition: transform 0.3s ease;
}

.doctor-card:hover {
  transform: translateY(-5px);
}

.doctor-image {
  flex-shrink: 0;
}

.doctor-photo {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #f0f0f0;
}

.doctor-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.doctor-name {
  color: #003449;
  font-size: 1.3rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

.doctor-position {
  color: #2c3e50;
  font-size: 1rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
}

.doctor-experience {
  color: #666;
  font-size: 0.9rem;
  margin-bottom: 1rem;
}

.appointment-btn {
  padding: 0.8rem 1.5rem;
  background: rgb(6, 113, 133);
  color: white;
  border: none;
  border-radius: 25px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
  align-self: flex-start;
}

.appointment-btn:hover {
  background: rgba(6, 113, 133, 0.9);
}

/* Стили для модального окна с iframe */
.iframe-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.8);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 10px;
  width: 60%;
  max-width: 800px;
  height: 90%;
  display: flex;
  flex-direction: column;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border-bottom: 1px solid #eee;
}

.modal-header h3 {
  margin: 0;
  color: #003449;
}

.close-btn {
  background: none;
  border: none;
  font-size: 2rem;
  cursor: pointer;
  color: #666;
  padding: 0;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.close-btn:hover {
  color: #000;
}

.appointment-iframe {
  width: 100%;
  height: 100%;
  border: none;
  border-radius: 0 0 10px 10px;
}
</style>
