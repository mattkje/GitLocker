<script setup lang="ts">
import { ref, onMounted } from 'vue'

async function getRepositoryDetails() {
  try {
    const response = await fetch('/api/repo-details/gitbay-core')
    if (!response.ok) throw new Error('Failed to fetch repository details')
    return await response.json()
  } catch (error) {
    console.error('Error fetching repository details:', error)
    return null
  }
}

const repoInfo = ref({
  name: '',
  owner: '',
  description: '',
  stars: 0,
  forks: 0,
  watchers: 0,
  updatedAt: '',
  defaultBranch: ''
})

const files = ref([])
const lastCommit = ref({
  message: '',
  hash: '',
  author: '',
  date: ''
})

onMounted(async () => {
  const data = await getRepositoryDetails()
  if (data) {
    repoInfo.value = {
      name: data.name,
      owner: data.owner,
      description: data.description,
      stars: data.stars,
      forks: data.forks,
      watchers: data.watchers,
      updatedAt: data.updatedAt,
      defaultBranch: data.defaultBranch
    }
    files.value = data.files || []
    lastCommit.value = data.lastCommit || { message: '', hash: '', author: '', date: '' }
  }
})
</script>

<template>
  <nav class="gh-tabs">
    <a class="active">Code</a>
    <a>Issues <span class="gh-tab-count">0</span></a>
    <a>Pull requests <span class="gh-tab-count">0</span></a>
    <a>Actions</a>
    <a>Projects</a>
    <a>Wiki</a>
    <a>Security</a>
    <a>Insights</a>
  </nav>
  <div class="gh-repo">
    <header class="gh-header">
      <div class="gh-header-title">
        <svg height="24" width="24" viewBox="0 0 16 16" class="gh-repo-icon">
          <path fill="currentColor"
                d="M2 2.5A2.5 2.5 0 0 1 4.5 0h7A2.5 2.5 0 0 1 14 2.5v11a.5.5 0 0 1-.757.429L8 10.101l-5.243 3.828A.5.5 0 0 1 2 13.5v-11Z"/>
        </svg>
        <span class="gh-owner">{{ repoInfo.owner }}</span>
        <span class="gh-slash">/</span>
        <span class="gh-name">{{ repoInfo.name }}</span>
        <span class="gh-public">Public</span>
      </div>
      <div class="gh-header-actions">
        <button class="gh-btn">Star <span>{{ repoInfo.stars }}</span></button>
        <button class="gh-btn">Fork <span>{{ repoInfo.forks }}</span></button>
        <button class="gh-btn">Watch <span>{{ repoInfo.watchers }}</span></button>
      </div>
    </header>
    <div class="gh-main">
      <div class="gh-main-col">
        <div class="gh-desc">{{ repoInfo.description }}</div>
        <div class="gh-branch-bar">
                    <span class="gh-branch">
                      <svg width="16" height="16" fill="none"><path fill="currentColor"
                                                                    d="M7.25 2.75a.75.75 0 0 1 .75-.75h.5a.75.75 0 0 1 0 1.5h-.5a.75.75 0 0 1-.75-.75Z"/></svg>
                      <b>{{ repoInfo.defaultBranch }}</b>
                    </span>
          <span class="gh-commit-msg">
                      <b>{{ lastCommit.message }}</b>
                      <span class="gh-commit-meta"> {{ lastCommit.hash }} by {{
                          lastCommit.author
                        }} on {{ lastCommit.date }}</span>
                    </span>
        </div>
        <table class="gh-file-list">
          <tbody>
          <tr v-for="file in files" :key="file.name">
            <td>
              <span v-if="file.type === 'dir'" class="gh-file-icon">
                <img src="@/assets/folder.svg" alt="" />
              </span>
              <span v-else class="gh-file-icon">
                <img src="@/assets/file.svg" alt="" />
              </span>
              <span class="gh-file-name">{{ file.name }}</span>
            </td>
            <td class="gh-file-msg"></td>
            <td class="gh-file-date"></td>
          </tr>
          </tbody>
        </table>
      </div>
      <aside class="gh-sidebar">
        <div class="gh-about">
          <h3>About</h3>
          <p>{{ repoInfo.description }}</p>
        </div>
        <ul class="gh-stats">
          <li>Stars: {{ repoInfo.stars }}</li>
          <li>Forks: {{ repoInfo.forks }}</li>
          <li>Watchers: {{ repoInfo.watchers }}</li>
          <li>Updated: {{ repoInfo.updatedAt }}</li>
        </ul>
      </aside>
    </div>
  </div>
</template>

<style scoped>
.gh-repo {
  font-family: system-ui, sans-serif;
  max-width: 1100px;
  margin: 2rem auto;
}

.gh-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.2rem 2rem 1rem 2rem;
  border-bottom: 1px solid #d0d7de;
}

.gh-header-title {
  display: flex;
  align-items: center;
  font-size: 1.3rem;
  gap: 0.3rem;
}

.gh-repo-icon {
  margin-right: 0.5rem;
  color: #57606a;
}

.gh-owner, .gh-name {
  font-weight: 600;
}

.gh-slash {
  color: #6e7781;
}

.gh-public {
  background: #ddf4ff;
  color: #0969da;
  font-size: 0.85rem;
  border-radius: 2em;
  padding: 0.1em 0.7em;
  margin-left: 0.7em;
  font-weight: 500;
}

.gh-header-actions {
  display: flex;
  gap: 0.5rem;
}

.gh-btn {
  background: #f6f8fa;
  border: 1px solid #d0d7de;
  border-radius: 6px;
  padding: 0.3em 1em;
  font-size: 1rem;
  cursor: pointer;
  color: #24292f;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 0.3em;
}

.gh-tabs {
  display: flex;
  gap: 2rem;
  padding: 0.5rem 2rem;
  border-bottom: 1px solid #d0d7de;
  background: #f6f8fa;
  font-size: 0.8rem;
}

.gh-tabs a {
  color: #57606a;
  text-decoration: none;
  padding: 0.3em 0.5em;
  border-bottom: 2px solid transparent;
  cursor: pointer;
}

.gh-tabs a.active {
  color: #24292f;
  border-bottom: 2px solid #fd8c73;
  font-weight: 600;
}

.gh-tab-count {
  background: #eaeef2;
  border-radius: 2em;
  padding: 0 0.5em;
  font-size: 0.9em;
  margin-left: 0.2em;
}

.gh-main {
  display: flex;
  flex-direction: row;
  padding: 2rem;
  gap: 2rem;
}

.gh-main-col {
  flex: 1.7;
}

.gh-desc {
  font-size: 1.1rem;
  color: #57606a;
  margin-bottom: 1.2rem;
}

.gh-branch-bar {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  background: #f6f8fa;
  border: 1px solid #d0d7de;
  border-radius: 6px;
  padding: 0.7em 1em;
  margin-bottom: 1.2rem;
  font-size: 1rem;
}

.gh-branch svg {
  vertical-align: middle;
  margin-right: 0.3em;
}

.gh-commit-msg {
  color: #57606a;
}

.gh-commit-meta {
  color: #8c959f;
  font-size: 0.95em;
  margin-left: 0.5em;
}

.gh-file-list {
  border: 1px solid #d0d7de;
  border-radius: 6px;
  padding: 1rem;
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 2rem;
}

.gh-file-list td {
  padding: 0.5em 0.2em;
  font-size: 0.8rem;
  border-bottom: 1px solid #f6f8fa;
}

.gh-file-icon {
  margin-right: 0.5em;
}
.gh-file-icon img {
  width: 18px;
  height: 18px;
}

.gh-file-name {
  color: #0969da;
  cursor: pointer;
}

.gh-sidebar {
  flex: 1;
  max-width: 270px;
}

.gh-about h3 {
  margin-top: 0;
  font-size: 1.1rem;
}

.gh-about p {
  color: #57606a;
  font-size: 1rem;
}

.gh-stats {
  list-style: none;
  padding: 0;
  margin: 1.5em 0 0 0;
  color: #57606a;
  font-size: 0.98rem;
}

.gh-stats li {
  margin-bottom: 0.5em;
}
</style>