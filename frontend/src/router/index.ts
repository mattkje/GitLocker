import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import RepositoriesPage from '@/views/RepositoriesPage.vue'

const routes: RouteRecordRaw[] = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/repos', name: 'Repositories', component: RepositoriesPage },
  // Add more routes here
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router