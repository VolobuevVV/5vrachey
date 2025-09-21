import { createRouter, createWebHistory } from 'vue-router'
import DepartmentPage from '@/components/DepartmentPage.vue'
import MainPage from '@/components/MainPage.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: MainPage
  },
  {
    path: '/department/:id',
    name: 'Department',
    component: DepartmentPage,
    props: true
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
