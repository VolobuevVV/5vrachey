<template>
  <div class="doctors-list">
    <h2 class="doctors-title">Наши врачи</h2>

    <!-- Фильтры -->
    <div class="filters-section">
      <div class="filter-group">
        <label class="filter-label">Отделение:</label>
        <select v-model="selectedDepartment" class="filter-select">
          <option value="">Все отделения</option>
          <option
            v-for="department in departments"
            :key="department.id"
            :value="department.id"
          >
            {{ department.name }}
          </option>
        </select>
      </div>

      <div class="filter-group">
        <label class="filter-label">Должность:</label>
        <select v-model="selectedPosition" class="filter-select">
          <option value="">Все должности</option>
          <option
            v-for="position in uniquePositions"
            :key="position"
            :value="position"
          >
            {{ position }}
          </option>
        </select>
      </div>
    </div>

    <!-- Сетка врачей -->
    <div class="doctors-grid">
      <DoctorCard
        v-for="doctor in filteredDoctors"
        :key="doctor.id"
        :doctor="doctor"
      />
    </div>

    <!-- Сообщение если нет врачей -->
    <div v-if="filteredDoctors.length === 0" class="no-doctors">
      <p>Врачи по выбранным фильтрам не найдены</p>
    </div>
  </div>
</template>

<script>
import DoctorCard from './DoctorCard.vue'
import { ref, computed, onMounted } from 'vue'

export default {
  name: 'DoctorsList',
  components: {
    DoctorCard
  },
  setup() {
    const doctors = ref([])
    const departments = ref([])
    const selectedDepartment = ref('')
    const selectedPosition = ref('')

    // Загрузка врачей
    const loadDoctors = async () => {
      try {
        const response = await fetch('/api/doctors')
        doctors.value = await response.json()
      } catch (error) {
        console.error('Ошибка загрузки врачей:', error)
      }
    }

    // Загрузка отделений
    const loadDepartments = async () => {
      try {
        const response = await fetch('/api/departments')
        departments.value = await response.json()
      } catch (error) {
        console.error('Ошибка загрузки отделений:', error)
      }
    }

    // Уникальные должности
    const uniquePositions = computed(() => {
      const positions = new Set()
      doctors.value.forEach(doctor => {
        doctor.positions.forEach(position => positions.add(position))
      })
      return Array.from(positions).sort()
    })

    // Отфильтрованные врачи
    const filteredDoctors = computed(() => {
      let filtered = doctors.value

      if (selectedDepartment.value) {
        filtered = filtered.filter(doctor =>
          doctor.departments.includes(selectedDepartment.value)
        )
      }

      if (selectedPosition.value) {
        filtered = filtered.filter(doctor =>
          doctor.positions.includes(selectedPosition.value)
        )
      }

      return filtered
    })

    onMounted(() => {
      loadDoctors()
      loadDepartments()
    })

    return {
      doctors,
      departments,
      selectedDepartment,
      selectedPosition,
      uniquePositions,
      filteredDoctors
    }
  }
}
</script>

<style scoped>
.doctors-list {
  width: 100%;
}

.doctors-title {
  color: rgb(0, 52, 73);
  font-size: 1.8rem;
  margin-bottom: 2rem;
  font-weight: 700;
  text-align: center;
}

.filters-section {
  display: flex;
  gap: 2rem;
  margin-bottom: 2rem;
  flex-wrap: wrap;
  justify-content: center;
}

.filter-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.filter-label {
  font-weight: 600;
  color: #003449;
  font-size: 0.9rem;
}

.filter-select {
  padding: 0.5rem 1rem;
  border: 2px solid #e0e0e0;
  border-radius: 8px;
  font-size: 1rem;
  min-width: 200px;
  background: white;
  cursor: pointer;
}

.filter-select:focus {
  outline: none;
  border-color: rgb(6, 113, 133);
}

.doctors-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
  justify-items: center;
}

.no-doctors {
  text-align: center;
  padding: 3rem;
  color: #666;
  font-size: 1.1rem;
}

</style>
