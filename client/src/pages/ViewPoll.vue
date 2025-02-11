<script setup>
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const pollId = route.params.id;
const pollData = ref(null);

onMounted(async () => {
  try {
    const response = await fetch(`http://localhost:8080/poll/${pollId}`);
    const result = await response.json();
    pollData.value = result.data;
    console.log(pollData.value);
  } catch (error) {
    console.log("Error fetching poll");
  }
});
</script>

<template>
  <div class="flex flex-col items-center min-h-screen w-full bg-sky-800 py-10">
    <div
      v-if="pollData !== null"
      class="border p-6 rounded-lg shadow-lg bg-white bg-opacity-80 max-w-xl w-full"
    >
      <h1 class="font-bold text-xl">{{ pollData?.title }}</h1>
      <p class="text-gray-700">{{ pollData?.description }}</p>
      <div class="mt-4">
        <h2 class="font-semibold">Make a choice:</h2>
        <ul class="mt-2">
          <li
            v-for="(choice, index) in pollData?.options.split('|')"
            :key="index"
            class="p-2"
          >
            <div class="">
              <input class="" type="checkbox" /><label
                :for="'checkbox-' + index"
                class="ml-2"
                >{{ choice }}</label
              >
            </div>
          </li>
        </ul>
      </div>
    </div>
    <div class="text-white font-bold text-3xl" v-else>Poll not found</div>
  </div>
</template>
