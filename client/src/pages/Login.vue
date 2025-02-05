<script setup>
import { ref } from "vue";
import { createAccount, login } from "../services/LoginService";
import InputField from "../components/InputField.vue";
import router from "../router";

const username = ref("");
const password = ref("");
const infoText = ref("");
const registerMode = ref(false);
const isSuccess = ref(false);

const reset = () => {
  username.value = "";
  password.value = "";
  infoText.value = "";
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
      if (!registerMode.value) {
        router.push({ name: "Main" });
      }
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
  <div class="flex items-center justify-center h-screen w-full bg-sky-800">
    <form
      @submit.prevent="handleData"
      class="border p-6 rounded-lg shadow-lg bg-white bg-opacity-80 max-w-sm w-full"
    >
      <h2 class="text-2xl font-bold mb-4 text-center text-gray-800">
        {{ registerMode ? "Register" : "Login" }}
      </h2>

      <InputField
        type="text"
        id="username"
        label="Username"
        v-model="username"
        placeholder="Enter your username"
      />

      <InputField
        type="password"
        id="password"
        label="Password"
        v-model="password"
        placeholder="Enter your password"
      />

      <div class="flex items-center justify-center mt-3">
        <button
          type="submit"
          class="bg-blue-600 text-white rounded w-full h-10 font-bold hover:bg-blue-700 transition duration-300 disabled:bg-gray-400 disabled:cursor-not-allowed"
          :disabled="username === '' || password === ''"
        >
          <div>{{ registerMode ? "Register" : "Login" }}</div>
        </button>
      </div>
      <div class="flex items-center justify-center pt-2">
        <div v-if="!registerMode">
          Don't have an account?
          <span
            class="hover:underline cursor-pointer text-blue-500"
            @click="(registerMode = true), reset()"
            >register</span
          >
        </div>
        <div v-if="registerMode">
          Already have an account?
          <span
            class="hover:underline cursor-pointer text-blue-500"
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
