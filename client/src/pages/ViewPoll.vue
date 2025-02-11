<script setup>
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const pollId = route.params.id;
const pollData = ref(null);

onMounted(async () => {
  try {
    const response = await fetch(`http://localhost:8080/poll/${pollId}`);
    pollData.value = await response.json();
    console.log(pollData.value);
  } catch (error) {
    console.log("Error fetching poll");
  }
});
</script>

<template>
  <div class="min-h-screen w-full bg-sky-800">
    <h1>Poll Details</h1>
    <pre>{{ pollData }}</pre>
  </div>
</template>
