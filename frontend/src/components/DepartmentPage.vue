<template>
  <div class="department-page">
    <div class="page-header">
      <h1 class="department-name">{{ departmentName }}</h1>
    </div>

    <div class="doctors-grid">
      <DoctorCard
        v-for="doctor in doctors"
        :key="doctor.id"
        :doctor="doctor"
      />
    </div>
  </div>
</template>

<script>
import DoctorCard from '@/components/DoctorCard.vue'

export default {
  name: 'DepartmentPage',
  components: {
    DoctorCard
  },
  data() {
    return {
      doctors: [],
      departmentName: ''
    }
  },
  computed: {
    departmentId() {
      return this.$route.params.id
    }
  },
  async mounted() {
    await this.loadDoctors()
    await this.loadDepartmentName()
  },
  methods: {
    async loadDoctors() {
      try {
        const response = await fetch(`/api/doctors/department/${this.departmentId}`)
        this.doctors = await response.json()
      } catch (error) {
        console.error('Ошибка загрузки врачей:', error)
      }
    },
    async loadDepartmentName() {
      try {
        const response = await fetch('/api/departments')
        const departments = await response.json()
        const department = departments.find(dept => dept.id === this.departmentId)
        this.departmentName = department ? department.full_name : 'Отделение'
      } catch (error) {
        console.error('Ошибка загрузки названия отделения:', error)
        this.departmentName = 'Отделение'
      }
    }
  }
}
</script>

<style scoped>
.department-page {
  padding-top: 160px;
  width: 100%;
  height: 100vh;
  background-color: rgb(255, 255, 255);
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
}

.page-header {
  text-align: center;
  width: 100%;
}

.department-name {
  color: #003449;
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 30px;
}

.doctors-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, 300px);
  gap: 2rem;
  width: 100%;
  justify-content: center;
}
</style>
