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

        <button
          class="appointment-btn"
          @click="openAppointment"
          v-if="doctor.available_for_appointment"
        >
          Записаться онлайн
        </button>
      </div>
    </div>

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
.doctor-card {
  background: white;
  border-radius: 15px;
  padding: 1.5rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.5rem;
  transition: transform 0.3s ease;
  text-align: center;
  height: 100%;
}

.doctor-card:hover {
  transform: translateY(-5px);
}

.doctor-image {
  display: flex;
  justify-content: center;
  width: 100%;
  flex-shrink: 0;
}

.doctor-photo {
  width: 120px;
  height: auto;
  aspect-ratio: 1;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid #f0f0f0;
}

.doctor-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.3rem;
  width: 100%;
  flex: 1;
  justify-content: space-between;
}

.doctor-name {
  color: #003449;
  font-size: 1.3rem;
  font-weight: 700;
  margin: 0;
  max-width: calc(100% - 60px);
  word-wrap: break-word;
  line-height: 1.3;
  margin-bottom: 0.2rem;
}

.doctor-position {
  color: #2c3e50;
  font-size: 1rem;
  font-weight: 600;
  margin: 0;
  max-width: calc(100% - 60px);
  word-wrap: break-word;
  line-height: 1.3;
  margin-bottom: 0.1rem;
}

.doctor-experience {
  color: #666;
  font-size: 0.9rem;
  margin: 0;
  margin-bottom: 0.5rem;
}

.appointment-btn {
  padding: 0.7rem 0;
  background: rgb(6, 113, 133);
  color: white;
  border: none;
  border-radius: 25px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
  width: 210px;
  font-size: 0.9rem;
  flex-shrink: 0;
}

.appointment-btn:hover {
  background: rgba(6, 113, 133, 0.9);
}

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
