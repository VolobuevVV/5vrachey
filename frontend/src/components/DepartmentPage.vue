<template>
  <div class="department-page">
    <div class="page-header">
      <h1 class="department-name">{{ departmentName }}</h1>
      <p class="department-description">Квалифицированные специалисты с многолетним опытом работы</p>
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
import doctorsData from '@/data/doctors.json'

export default {
  name: 'DepartmentPage',
  components: {
    DoctorCard
  },
  data() {
    return {
      doctors: []
    }
  },
  computed: {
    departmentId() {
      return this.$route.params.id
    },
    departmentName() {
      const department = doctorsData.departments.find(dept => dept.id === this.departmentId)
      return department ? department.name : 'Отделение'
    }
  },
  mounted() {
    this.loadDoctors()
  },
  methods: {
    loadDoctors() {
      const department = doctorsData.departments.find(dept => dept.id === this.departmentId)
      this.doctors = department ? department.doctors : []
    }
  }
}
</script>

<style scoped>
.department-page {
  padding-top: 30vh;
  max-width: 1200px;
}

.page-header {
  text-align: center;
  padding-top: 21vh;
}

.department-name {
  color: #003449;
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: 1rem;
}

.department-description {
  color: #2c3e50;
  font-size: 1.2rem;
}

.doctors-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 2rem;
  margin-top: 5vh;
}

</style>
