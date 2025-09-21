<template>
  <div>
    <div class="doctor-card">
      <div class="doctor-image">
        <img :src="doctor.photo" :alt="doctor.name" class="doctor-photo">
      </div>
      <div class="doctor-info">
        <h3 class="doctor-name">{{ doctor.name }}</h3>
        <p class="doctor-position">{{ doctor.position }}</p>
        <p class="doctor-experience">Опыт работы: {{ doctor.experience }}</p>
        <button class="appointment-btn" @click="showIframe = true">
          Записаться онлайн
        </button>
      </div>
    </div>

    <!-- Модальное окно с iframe -->
    <div v-if="showIframe" class="iframe-modal">
      <div class="modal-content">
        <div class="modal-header">
          <h3>Запись к {{ doctor.name }}</h3>
          <button class="close-btn" @click="showIframe = false">×</button>
        </div>
        <iframe
            :src="doctor.appointmentLink"
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
  methods: {
    openAppointment() {
      this.showIframe = true;
    }
  }
}
</script>

<style scoped>
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
