<script setup>
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import SingleChoice from "../components/ChoiceField.vue";

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
      <SingleChoice
        :type="pollData?.multiple"
        :options="pollData?.options"
        :multipleChoice="pollData?.multiple"
      />

      <button
        type="submit"
        class="bg-blue-600 text-white rounded w-full h-10 font-bold hover:bg-blue-700 transition duration-300 cursor-pointer mt-4"
      >
        Vote
      </button>
    </div>
    <div class="text-white font-bold text-3xl" v-else>Poll not found</div>
  </div>
</template>
