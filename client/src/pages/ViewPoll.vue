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
  <div
    class="flex flex-col items-center justify-center min-h-screen w-full bg-sky-800"
  >
    <form
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
      <div v-if="pollData?.requireNames === '1'">
        <h1 class="flex flex-row gap-1">
          <div class="font-semibold">Name</div>
          (Required)
        </h1>
        <input
          class="border rounded p-2 w-full mt-1 focus:outline-none focus:ring-2 focus:ring-blue-500"
          type="text"
        />
      </div>

      <div class="flex flex-row justify-between mt-4">
        <button
          type="submit"
          class="bg-blue-600 text-white rounded w-25 h-10 font-semibold hover:bg-blue-700 transition duration-300 cursor-pointer"
        >
          Vote
        </button>
        <button
          class="bg-blue-600 text-white rounded w-25 h-10 font-semibold hover:bg-blue-700 transition duration-300 cursor-pointer"
        >
          Show results
        </button>
      </div>
    </form>
    <div class="text-white font-bold text-3xl" v-else>Poll not found</div>
  </div>
</template>
