<script setup>
import { onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import router from "../router";
import ChoiceField from "../components/ChoiceField.vue";

const route = useRoute();
const pollId = route.params.id;
const pollData = ref(null);
const selectedOptions = ref([]);
const userName = ref("");
const missingOptions = ref(false);

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

const submitVote = async () => {
  if (selectedOptions.value.length === 0) {
    missingOptions.value = true;
    return;
  }
  try {
    const response = await fetch(`http://localhost:8080/poll/${pollId}/vote`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Option: selectedOptions.value.join(","),
        Username: userName.value,
      }),
    });
    if (response.ok) {
      missingOptions.value = false;
      router.push(`/poll/${pollId}/result`);
    }
  } catch (error) {
    console.error("Database error:", error);
  }
};
</script>

<template>
  <div
    class="flex flex-col items-center justify-center min-h-screen w-full bg-sky-800"
  >
    <form
      @submit.prevent="submitVote"
      v-if="pollData !== null"
      class="border p-6 rounded-lg shadow-lg bg-white bg-opacity-80 max-w-xl w-full"
    >
      <h1 class="font-bold text-xl">{{ pollData?.title }}</h1>
      <p class="text-gray-700">{{ pollData?.description }}</p>
      <ChoiceField
        :type="pollData?.multiple"
        :options="pollData?.votes"
        :multipleChoice="pollData?.multiple"
        v-model="selectedOptions"
      />
      <!-- <div v-if="pollData?.requireNames === '1'">
        <h1 class="flex flex-row gap-1">
          <div class="font-semibold">Name</div>
          (Required)
        </h1>
        <input
          class="border rounded p-2 w-full mt-1 focus:outline-none focus:ring-2 focus:ring-blue-500"
          type="text"
          v-model="userName"
        />
      </div> -->

      <div
        v-if="missingOptions === true"
        class="mt-4 p-4 rounded-lg shadow-md bg-red-200 text-red-700"
      >
        Please choose at least one option
      </div>

      <div class="flex flex-row justify-between mt-4">
        <button
          type="submit"
          class="bg-blue-600 text-white rounded w-25 h-10 font-semibold hover:bg-blue-700 transition duration-300 cursor-pointer"
        >
          Vote
        </button>
        <button
          @click.prevent="console.log(selectedOptions)"
          class="bg-blue-600 text-white rounded w-25 h-10 font-semibold hover:bg-blue-700 transition duration-300 cursor-pointer"
        >
          Show results
        </button>
      </div>
    </form>
    <div class="text-white font-bold text-3xl" v-else>Poll not found</div>
  </div>
</template>
