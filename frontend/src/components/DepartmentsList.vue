<template>
  <div class="departments-list">
    <h2 class="departments-title">Отделения медицинского центра</h2>
    <div class="departments-container">
      <div
        v-for="department in departments"
        :key="department.id"
        class="department-item"
        @click="goToDepartment(department.id)"
      >
        <div class="department-icon">
          <img :src="getDepartmentIcon(department.id)" :alt="department.name">
        </div>
        <h3 class="department-name">{{ department.name }}</h3>
        <p class="department-description">Современное оборудование и квалифицированные специалисты</p>
        <button class="select-button">Выбрать</button>
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
  name: 'DepartmentsList',
  setup() {
    const router = useRouter()
    const departments = ref([])

    const getDepartmentIcon = (departmentId) => {
      try {
        return new URL(`/src/assets/icons/${departmentId}.svg`, import.meta.url).href
      } catch (error) {
        return new URL(`/src/assets/icons/default.svg`, import.meta.url).href
      }
    }

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
      goToDepartment,
      getDepartmentIcon
    }
  }
}
</script>

<style scoped>
.departments-list {
  width: 100%;
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
  grid-template-columns: repeat(4, 1fr);
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.department-item {
  background: white;
  border: 2px solid rgb(6, 113, 133);
  border-radius: 15px;
  padding: 2rem;
  transition: all 0.3s ease;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
  min-height: 280px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.department-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 15px rgba(0, 0, 0, 0.15);
}

.department-icon {
  width: 60px;
  height: 60px;
  margin-bottom: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.department-icon img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.department-name {
  color: rgb(0, 52, 73);
  font-size: 1.3rem;
  font-weight: 700;
  margin-bottom: 1rem;
  line-height: 1.3;
}

.department-description {
  color: #666;
  font-size: 0.9rem;
  margin-bottom: 1.5rem;
  line-height: 1.4;
  flex-grow: 1;
}

.select-button {
  background: rgb(6, 113, 133);
  color: white;
  border: none;
  border-radius: 25px;
  padding: 0.8rem 2rem;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
  width: 160px;
}

.select-button:hover {
  background: rgba(6, 113, 133, 0.9);
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
