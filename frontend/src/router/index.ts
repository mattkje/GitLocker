import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomePage from '@/views/HomePage.vue'

const routes: RouteRecordRaw[] = [
  { path: '/', name: 'Home', component: HomePage },
  // Add more routes here
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router