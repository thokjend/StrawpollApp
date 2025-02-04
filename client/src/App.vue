<script setup>
import { ref } from "vue";

const username = ref("");
const password = ref("");
const registerMode = ref(false);

const login = () => {
  console.log(username.value);
};
</script>

<template>
  <div class="flex items-center justify-center h-screen w-full">
    <form @submit.prevent="login" class="border p-4 rounded-lg shadow-lg">
      <div class="flex items-center">
        <label for="username" class="block text-lg font-bold pr-3"
          >Username:
        </label>
        <input
          type="text"
          class="border rounded p-2 w-full mt-2"
          v-model="username"
        />
      </div>

      <div class="flex items-center">
        <label for="password" class="block text-lg font-bold pr-3"
          >Password:
        </label>
        <input
          type="password"
          class="border rounded p-2 w-full mt-2"
          v-model="password"
        />
      </div>

      <div class="flex items-center justify-center mt-3">
        <button
          type="submit"
          class="bg-sky-500 border rounded w-full h-10 font-bold"
        >
          Login
        </button>
      </div>
    </form>
  </div>
</template>
