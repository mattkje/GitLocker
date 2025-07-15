<script setup>
import {ref} from "vue";
import CreateRepo from "@/components/CreateRepo.vue";

// Mock user state; replace with real auth logic
const user = ref(null); // or { name: "Matt", avatar: "/avatar.png" }

const showModal = ref(false);

function openModal() {
  showModal.value = true;
}

function closeModal() {
  showModal.value = false;
}

function login() {
  // Replace with real login logic
  user.value = {name: "Matt", avatar: "src/assets/profile.png"};
}

function logout() {
  user.value = null;
}
</script>

<template>
  <header>
    <div class="left-content">
      <div class="wrapper">
        <img class="logoImg" src="@/assets/logo.png" alt="GitLocker"/>
        <h1 class="logo">GitBay</h1>
      </div>
      <nav>
        <router-link class="tab" to="/">Home</router-link>
        <router-link class="tab" to="/repos">Repositories</router-link>
        <router-link class="tab" to="/settings">Settings</router-link>
      </nav>
    </div>
    <div class="right-content">
      <button id="createBtn" title="Create Repository" @click="openModal">＋</button>
      <div class="profile">
        <template v-if="user">
          <img :src="user.avatar" alt="Profile" class="avatar" @click="logout" title="Logout"/>
        </template>
        <template v-else>
          <button class="loginBtn" @click="login">Login</button>
        </template>
      </div>
    </div>
  </header>
  <CreateRepo v-if="showModal" @close="closeModal"/>
</template>

<style scoped>
.logo {
  font-family: Lexend, sans-serif;
  letter-spacing: -1.36px;
  font-weight: 700;
  font-size: 1.2rem;
  user-select: none;
  color: var(--heading);
}

nav {
  margin-left: 20px;
  display: flex;
  gap: 20px;
  font-weight: 600;
}

nav .tab {
  background: none;
  border: none;
  color: var(--text);
  font-size: 0.9rem;
  cursor: pointer;
  padding: 5px 5px;
  border-radius: 6px;
  transition: background-color 0.2s ease;
  text-decoration: none;
  outline: none;
}

nav .tab.router-link-active,
nav .tab:hover {
  background: var(--border);
  color: var(--heading);
}

.logoImg {
  width: 32px;
  height: 32px;
}

.wrapper {
  display: flex;
  align-items: center;
}

header {
  background: var(--background-card);
  color: var(--heading);
  padding: 5px 20px;
  border-bottom: 1px solid var(--border);
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: sticky;
  top: 0;
  z-index: 1000;
}

#createBtn {
  background: none;
  border: none;
  font-size: 1.6rem;
  cursor: pointer;
  color: var(--text);
  padding: 5px 10px;
  border-radius: 6px;
  transition: background-color 0.2s ease;
}

#createBtn:hover {
  color: var(--accent);
}

.profile {
  display: flex;
  align-items: center;
  margin-left: 16px;
}

.avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  cursor: pointer;
  border: 2px solid var(--accent);
  transition: box-shadow 0.2s;
}

.avatar:hover {
  box-shadow: 0 0 0 2px var(--accent);
}

.loginBtn {
  background: none;
  color: var(--text);
  border: 0.5px solid var(--text);
  border-radius: 6px;
  padding: 6px 14px;
  font-size: 0.8rem;
  cursor: pointer;
  transition: border-color 0.2s, color 0.2s;
}

.loginBtn:hover {
  border-color: var(--accent-hover);
  color: var(--accent-hover);
}
.left-content {
  display: flex;
  align-items: center;
}

.right-content {
  display: flex;
  align-items: center;
}
</style>