<template>
  <div class="departments-grid">
    <h2 class="departments-title">Отделения медицинского центра</h2>
    <div class="departments-container">
      <div
        v-for="department in departments"
        :key="department.id"
        class="department-item"
        :style="{ backgroundColor: department.color }"
        @click="goToDepartment(department.id)"
      >
        {{ department.name }}
      </div>
    </div>
    <div class="hero-widget">
      <img src="@/components/icons/medal.svg" class="medal-icon" alt="Медаль">
      <span class="widget-text">Входим в топ 3 частных клиник Адлера на ПроДокторов</span>
    </div>
  </div>
</template>

<script>
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'

export default {
  name: 'DepartmentsGrid',
  setup() {
    const router = useRouter()
    const departments = ref([])

    const goToDepartment = (departmentId) => {
      router.push(`/department/${departmentId}`)
    }

    const loadDepartments = async () => {
      try {
        const response = await fetch('/api/departments')
        departments.value = await response.json()
      } catch (error) {
        console.error('Ошибка загрузки отделений:', error)
      }
    }

    onMounted(() => {
      loadDepartments()
    })

    return {
      departments,
      goToDepartment
    }
  }
}
</script>

<style scoped>


.department-item {
  cursor: pointer;
}


.departments-grid {
  width: 90%;
  height: 100%;
}

.departments-title {
  color: rgb(0, 52, 73);
  font-size: 1.8rem;
  margin-bottom: 2rem;
  font-weight: 700;
  text-align: center;
}

.departments-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 0.8rem;
  margin-bottom: 2rem;
}

.department-item {
  padding: 1rem;
  border-radius: 8px;
  transition: all 0.3s ease;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  color: rgb(0, 52, 73);
  font-size: 1.1rem;
  min-height: 60px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.department-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.hero-widget {
  padding: 0.8rem 1.5rem;
  background: rgba(213, 238, 241, 0.3);
  color: #2c3e50;
  border-radius: 15px;
  font-size: 1.1rem;
  font-weight: 600;
  text-align: center;
  width: 100%;
  grid-column: 1 / -1;
  margin-top: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.medal-icon {
  width: 40px;
  height: 40px;
}

.widget-text {
  font-size: 1.1rem;
  font-weight: 600;
  color: #003449;
}
</style>
