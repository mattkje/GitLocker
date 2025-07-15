<template>
  <div class="modal-backdrop" @click.self="close">
    <div class="modal-content" role="dialog" aria-modal="true" aria-labelledby="modalTitle">
      <button class="closeBtn" aria-label="Close modal" @click="close">&times;</button>
      <h3 id="modalTitle">Create a New Repository</h3>
      <form id="repoForm" @submit.prevent="submit">
        <input
          type="text"
          id="repoName"
          name="repoName"
          v-model="repoName"
          placeholder="Repository name"
          required
          autocomplete="off"
          minlength="1"
          maxlength="100"
        />
        <button type="submit">Create Repository</button>
      </form>
      <div id="status">{{ status }}</div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const emit = defineEmits(['close', 'create'])

const repoName = ref('')
const status = ref('')

function close() {
  emit('close')
}

async function submit() {
  const name = repoName.value.trim()
  status.value = ''
  if (!name) {
    status.value = 'Repository name is required.'
    return
  }
  try {
    const res = await fetch('/api/create-repo', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name }),
    })
    if (res.ok) {
      status.value = `Repository "${name}" created successfully!`
      emit('create', name)
      repoName.value = ''
    } else {
      const errText = await res.text()
      status.value = `Error: ${errText}`
    }
  } catch (err) {
    status.value = `Error: ${err.message}`
  }
  // Optionally close modal on success
  // if (res.ok) close()
}
</script>

<style scoped>
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
}

.modal-content {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  min-width: 300px;
  position: relative;
}

.closeBtn {
  position: absolute;
  top: 10px;
  right: 14px;
  background: none;
  border: none;
  font-size: 1.6rem;
  cursor: pointer;
  color: #333;
}


.modal-content{
  background: white;
  border-radius: 8px;
  padding: 25px 30px;
  box-shadow: 0 8px 24px rgb(0 0 0 / 0.2);
  width: 90%;
  max-width: 400px;
  position: relative;
}
.modal-content h3 {
  margin-top: 0;
  font-weight: 700;
  color: #24292e;
}
.modal-content form {
  display: flex;
  flex-direction: column;
  gap: 15px;
  margin-top: 15px;
}
.modal-content input[type="text"] {
  padding: 10px;
  font-size: 1rem;
  border: 1px solid #d1d5da;
  border-radius: 6px;
  outline-offset: 2px;
}
.modal-content input[type="text"]:focus {
  border-color: #0366d6;
  box-shadow: 0 0 0 3px rgb(3 102 214 / 0.3);
}
.modal-content button {
  padding: 10px;
  font-size: 1rem;
  background-color: #2ea44f;
  border: none;
  color: white;
  font-weight: 600;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}
.modal-content button:hover {
  background-color: #22863a;
}
.modal-content .closeBtn {
  position: absolute;
  top: 12px;
  right: 15px;
  font-size: 1.4rem;
  border: none;
  background: none;
  cursor: pointer;
  color: #666;
  font-weight: 700;
}
.modal-content .closeBtn:hover {
  color: #24292e;
}
#status {
  margin-top: 12px;
  font-weight: 600;
}
#status.error {
  color: #d73a49;
}
#status.success {
  color: #28a745;
}
</style>