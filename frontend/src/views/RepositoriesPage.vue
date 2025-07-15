<script setup lang="ts">
import { ref, computed } from 'vue'

interface Repo {
  name: string
  description: string
  updatedAt: string
}

const repos = ref<Repo[]>([
  {
    name: 'gitbay-core',
    description: 'Core backend for GitBay',
    updatedAt: '2024-06-01',
  },
  {
    name: 'gitbay-frontend',
    description: 'Vue.js frontend for GitBay',
    updatedAt: '2024-05-28',
  },
  {
    name: 'gitbay-utils',
    description: 'Utility scripts and helpers',
    updatedAt: '2024-05-20',
  },
])

const search = ref('')

const filteredRepos = computed(() =>
  repos.value.filter(repo =>
    repo.name.toLowerCase().includes(search.value.toLowerCase())
  )
)
</script>

<template>
  <div class="repositories-page">
    <h1>Repositories</h1>
    <div class="search-bar">
      <input
        v-model="search"
        type="text"
        placeholder="Find a repository..."
      />
    </div>
    <ul class="repo-list">
      <li v-for="repo in filteredRepos" :key="repo.name" class="repo-item">
        <div class="repo-header">
          <a class="repo-name" href="#">{{ repo.name }}</a>
          <span class="repo-updated">Updated on {{ repo.updatedAt }}</span>
        </div>
        <div class="repo-desc">{{ repo.description }}</div>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.repositories-page {
  width: 100%;
  min-height: 100vh;
  background: var(--background);
}

h1 {
  font-size: 1.7rem;
  margin: 2rem 0 1.2rem 0;
  color: var(--heading);
  padding-left: 2rem;
}

.search-bar {
  margin-bottom: 1.5rem;
  padding: 0 2rem;
}

.search-bar input {
  width: 100%;
  padding: 8px 12px;
  border-radius: 6px;
  border: 1px solid var(--border);
  background: var(--background);
  color: var(--text);
  font-size: 1rem;
  transition: border-color 0.2s;
}

.search-bar input:focus {
  outline: none;
  border-color: var(--accent);
}

.repo-list {
  list-style: none;
  padding: 0 2rem;
  margin: 0;
}

.repo-item {
  padding: 18px 0 14px 0;
  border-bottom: 1px solid var(--border);
}

.repo-item:last-child {
  border-bottom: none;
}

.repo-header {
  display: flex;
  align-items: center;
  gap: 12px;
}

.repo-name {
  font-weight: 600;
  color: var(--accent);
  text-decoration: none;
  font-size: 1.1rem;
}

.repo-name:hover {
  text-decoration: underline;
  color: var(--accent-hover);
}

.repo-updated {
  color: var(--text);
  font-size: 0.95em;
}

.repo-desc {
  color: var(--text);
  margin-top: 4px;
  font-size: 0.98em;
}
</style>