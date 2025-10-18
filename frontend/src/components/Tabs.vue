<template>
  <div class="tabs-container">
    <div class="tabs-header">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        :class="['tab-button', { 'active': activeTab === tab.id }]"
        @click="switchTab(tab.id)"
      >
        {{ tab.title }}
      </button>
    </div>
    <div class="tabs-content">
      <component :is="activeComponent" />
    </div>
  </div>
</template>

<script>
import DoctorsList from './DoctorsList.vue'
import DepartmentsList from './DepartmentsList.vue'

export default {
  name: 'Tabs',
  components: {
    DoctorsList,
    DepartmentsList
  },
  data() {
    return {
      activeTab: 'doctors',
      tabs: [
        { id: 'doctors', title: 'Врачи', component: 'DoctorsList' },
        { id: 'departments', title: 'Отделения', component: 'DepartmentsList' }
      ]
    }
  },
  computed: {
    activeComponent() {
      const tab = this.tabs.find(t => t.id === this.activeTab)
      return tab ? tab.component : 'DoctorsList'
    }
  },
  methods: {
    switchTab(tabId) {
      this.activeTab = tabId
    }
  }
}
</script>

<style scoped>
.tabs-container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

.tabs-header {
  display: flex;
  border-bottom: 2px solid #e0e0e0;
  margin-bottom: 2rem;
}

.tab-button {
  padding: 1rem 2rem;
  background: none;
  border: none;
  font-size: 1.1rem;
  font-weight: 600;
  color: #666;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
}

.tab-button:hover {
  color: #007bff;
}

.tab-button.active {
  color: rgb(0, 52, 73);
}

.tab-button.active::after {
  content: '';
  position: absolute;
  bottom: -2px;
  left: 0;
  width: 100%;
  height: 3px;
  background-color: rgb(6, 113, 133);
}

.tabs-content {
  min-height: 400px;
}
</style>
