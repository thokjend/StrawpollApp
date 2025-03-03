<script setup>
import { onMounted, ref } from "vue";
import router from "../router";
import { ValidateToken } from "../utils/ValidateToken";

const username = ref("");

const logout = () => {};

onMounted(async () => {
  const token = await ValidateToken();
  username.value = token.username;
});
</script>

<template>
  <div class="flex flex-col items-center min-h-screen w-full bg-sky-800">
    <header
      class="text-white w-full h-15 flex items-center justify-around px-10 bg-sky-900"
    >
      <ul class="flex gap-10 font-bold text-xl">
        <li class="hover:underline cursor-pointer">Create Poll</li>
        <li class="hover:underline cursor-pointer">View Polls</li>
      </ul>

      <ul class="flex gap-5 font-bold text-xl">
        <li>{{ username }}</li>
        <li class="hover:underline cursor-pointer">Logout</li>
      </ul>
    </header>
  </div>
</template>
