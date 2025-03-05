<script setup>
import { onMounted, ref } from "vue";
import { ValidateToken } from "../utils/ValidateToken";
import Header from "../components/Header.vue";
import { fetchAllPolls, fetchSinglePoll } from "../services/PollService";
import { nextTick } from "vue";
import router from "../router";

const polls = ref([]);
const pollsDetails = ref({});

const navigateToPoll = (id) => {
  router.push(`/poll/${id}`);
};

onMounted(async () => {
  await fetchPollIds();
  await getPolls();
});

const getPolls = async () => {
  try {
    // Wait for all poll requests to finish
    const pollData = await Promise.all(
      polls.value.map(async (id) => {
        const response = await fetchSinglePoll(id);
        console.log(`Response for ${id}:`, response.data);
        return response.data;
      })
    );

    pollsDetails.value = pollData; // Store fetched poll details
  } catch (error) {
    console.log("Error fetching poll details", error);
  }
};

const fetchPollIds = async () => {
  try {
    const result = await fetchAllPolls();
    polls.value = result.polls;
  } catch (error) {
    console.log("Error fetching polls");
  }
};
</script>

<template>
  <div
    class="flex flex-col items-center min-h-screen w-full bg-sky-800 text-white p-30"
  >
    <Header></Header>
    <div class="w-full max-w-2xl">
      <h1 class="text-xl font-bold mb-4">Polls</h1>
      <div
        v-for="(poll, index) in pollsDetails"
        :key="index"
        class="mb-4 p-4 bg-gray-700 rounded-lg shadow-lg hover:bg-gray-800 cursor-pointer"
        @click="navigateToPoll(polls[index])"
      >
        <h2 class="text-lg font-semibold">{{ poll.title }}</h2>
        <p v-if="poll.description" class="text-sm text-gray-300">
          {{ poll.description }}
        </p>
        <div class="mt-2">
          <h3 class="text-md font-semibold">Votes:</h3>
          <ul>
            <li v-for="(voteCount, option) in poll.votes" :key="option">
              {{ option }}: {{ voteCount }}
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>
