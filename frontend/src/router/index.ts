import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import HomePage from '@/views/HomePage.vue'
import RepositoriesPage from '@/views/RepositoriesPage.vue'
import RepositoryPage from '@/views/RepositoryPage.vue'

const routes: RouteRecordRaw[] = [
  { path: '/', name: 'Home', component: HomePage },
  { path: '/repos', name: 'Repositories', component: RepositoriesPage },
  { path: '/repos/:id', name: 'Repositories', component: RepositoryPage, props: true }
  // Add more routes here
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router