<script setup>
import { onMounted, ref } from "vue";
import { ValidateToken } from "../utils/ValidateToken";
import router from "../router";

const username = ref("");

const logout = () => {
  localStorage.removeItem("sessionToken");
  router.push({ name: "Login" });
};
const navigateToCreate = () => {
  router.push({ name: "Create" });
};

const navigateToMain = () => {
  router.push({ name: "Main" });
};

onMounted(async () => {
  const token = await ValidateToken();
  if (token) {
    username.value = token.username;
  }
});
</script>

<template>
  <header
    class="fixed top-0 left-0 text-white w-full h-15 flex items-center justify-around px-10 bg-sky-900"
  >
    <ul class="flex gap-10 font-bold text-xl">
      <li class="hover:underline cursor-pointer" @click="navigateToCreate">
        Create Poll
      </li>
      <li class="hover:underline cursor-pointer" @click="navigateToMain">
        View Polls
      </li>
    </ul>

    <ul class="flex gap-5 font-bold text-xl">
      <li>{{ username }}</li>
      <li class="hover:underline cursor-pointer" @click="logout">Logout</li>
    </ul>
  </header>
</template>
