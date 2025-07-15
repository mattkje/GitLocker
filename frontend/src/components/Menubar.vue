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
  color: #cdd9e5;
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
  background: #0366d6;
  color: white;
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
  background: #24292e;
  color: white;
  padding: 5px 20px;
  border-bottom: 1px solid #cdd9e5;
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
  color: white;
  padding: 5px 10px;
  border-radius: 6px;
  transition: background-color 0.2s ease;
}

#createBtn:hover {
  background: #0366d6;
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
  border: 2px solid #0366d6;
  transition: box-shadow 0.2s;
}

.avatar:hover {
  box-shadow: 0 0 0 2px #0366d6;
}

.loginBtn {
  background: #2d8cf0;
  color: #fff;
  border: none;
  border-radius: 6px;
  padding: 6px 16px;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.2s;
}

.loginBtn:hover {
  background: #1a6fb3;
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