<script setup>
import { onMounted, ref } from "vue";
import router from "../router";
import { stringify } from "postcss";

defineProps({
  title: String,
  description: String,
  options: String,
  settings: Array,
});

onMounted(async () => {
  try {
    const response = await fetch("http://localhost:8080/poll/Test");
    const data = await response.json();
    console.log(data);
  } catch (error) {
    console.log("Error fetching poll");
  }
});
</script>

<template>
  <div class="min-h-screen w-full bg-sky-800"></div>
</template>
