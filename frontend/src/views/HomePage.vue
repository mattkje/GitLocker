<template>
  <div class="dashboard">
    <h1>Welcome, {{ user?.name || "Guest" }}</h1>
    <div class="commit-history">
      <h2>Commit History</h2>
      <div class="grid">
        <div
            v-for="week in 52"
            :key="week"
            class="week"
        >
          <div
              v-for="day in 7"
              :key="day"
              class="day"
              :style="{ background: getColor(commitGrid[week - 1][day - 1]) }"
              :title="`${commitGrid[week - 1][day - 1]} commits`"
          ></div>
        </div>
      </div>
    </div>
    <div class="recent">
      <h2>Recent Updates</h2>
      <ul>
        <li v-for="update in recentUpdates" :key="update.id">
          <strong>{{ update.repo }}</strong>: {{ update.message }} <span class="date">{{ update.date }}</span>
        </li>
        <li v-if="recentUpdates.length === 0">No recent updates.</li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import {ref, computed} from 'vue'

// Mock user
const user = ref({name: "Matt"})

// Flat array: 52 weeks * 7 days = 364 days
const commitData = ref(Array.from({length: 364}, () => Math.floor(Math.random() * 6)))

// Convert flat array to 52x7 grid
const commitGrid = computed(() => {
  const grid = []
  for (let w = 0; w < 52; w++) {
    grid.push(commitData.value.slice(w * 7, w * 7 + 7))
  }
  return grid
})

// Blue color scale
function getColor(count) {
  if (count === 0) return "#e3eafc"
  if (count < 2) return "#b6cdfa"
  if (count < 4) return "#7da7f7"
  if (count < 6) return "#3773e0"
  return "#174ea6"
}

// Mock recent updates
const recentUpdates = ref([
  {id: 1, repo: "my-first-repo", message: "Fixed bug in API", date: "2024-06-10"},
  {id: 2, repo: "project-x", message: "Added README", date: "2024-06-09"},
  {id: 3, repo: "test-repo", message: "Initial commit", date: "2024-06-08"}
])
</script>

<style scoped>
.dashboard {
  width: 100%;
  min-height: 100vh;
  background: var(--background);
  padding: 0;
  margin: 0;
  border-radius: 0;
  box-shadow: none;
  border: none;
}

h1 {
  font-size: 2rem;
  margin: 2rem 0 1.2rem 0;
  color: var(--heading);
  padding-left: 2rem;
}

.commit-history {
  margin-bottom: 2rem;
  padding-left: 2rem;
}

.grid {
  display: flex;
  gap: 2px;
}

.week {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.day {
  width: 14px;
  height: 14px;
  border-radius: 3px;
  border: 1px solid var(--border);
  transition: background 0.2s;
}

.recent {
  padding-left: 2rem;
}

.recent h2 {
  margin-bottom: 0.7rem;
  color: var(--heading);
}

ul {
  list-style: none;
  padding: 0;
  background: none;
  border-radius: 0;
  border: none;
}

li {
  margin-bottom: 0.5rem;
  color: var(--text);
  padding: 6px 0;
  border-bottom: 1px solid var(--border);
}

li:last-child {
  border-bottom: none;
}

.date {
  color: var(--accent);
  font-size: 0.9em;
  margin-left: 8px;
}
</style>