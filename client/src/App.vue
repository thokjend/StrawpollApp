<script setup>
import { ref } from "vue";
import { createAccount, login } from "./services/LoginService";

const username = ref("");
const password = ref("");
const infoText = ref("");
const registerMode = ref(false);
const isSuccess = ref(false);

const reset = () => {
  username.value = "";
  password.value = "";
};

const handleData = async () => {
  try {
    let response;
    if (registerMode.value) {
      response = await createAccount(username.value, password.value);
    } else {
      response = await login(username.value, password.value);
    }

    if (response.success) {
      isSuccess.value = true;
      infoText.value = registerMode.value ? "Success! Account created" : "";
    } else {
      isSuccess.value = false;
      infoText.value = registerMode.value
        ? "Account creation failed. User already exist."
        : "Login failed. Invalid username or password.";
    }
  } catch (error) {
    console.error("Error", error);
    infoText.value = "An unexpected error occurred. Please try again.";
  }
};
</script>

<template>
  <div class="flex items-center justify-center h-screen w-full">
    <form @submit.prevent="handleData" class="border p-4 rounded-lg shadow-lg">
      <div class="flex items-center">
        <label for="username" class="block text-lg font-bold pr-3 w-24"
          >Username:
        </label>
        <input
          type="text"
          class="border rounded p-2 w-full mt-2"
          v-model="username"
        />
      </div>

      <div class="flex items-center">
        <label for="password" class="block text-lg font-bold pr-3 w-24"
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
          class="bg-sky-500 border rounded w-full h-10 font-bold cursor-pointer disabled:bg-gray-400 disabled:cursor-default disabled:border-gray-300"
          :disabled="username === '' || password === ''"
        >
          <div v-if="registerMode == false">Login</div>
          <div v-if="registerMode == true">Register</div>
        </button>
      </div>
      <div class="flex items-center justify-center pt-2">
        <div v-if="registerMode == false">
          Don't have an account?
          <span
            class="hover:underline cursor-pointer"
            @click="(registerMode = true), reset()"
            >register</span
          >
        </div>
        <div v-if="registerMode == true">
          Already have an account?
          <span
            class="hover:underline cursor-pointer"
            @click="(registerMode = false), reset()"
            >Login</span
          >
        </div>
      </div>
      <div
        class="flex items-center justify-center pt-3 font-bold"
        :class="isSuccess ? 'text-green-500' : 'text-red-500'"
        v-if="infoText"
      >
        {{ infoText }}
      </div>
    </form>
  </div>
</template>
