<script setup>
import { onMounted, ref } from "vue";
import { ValidateToken } from "../utils/ValidateToken";
import Header from "../components/Header.vue";
import { fetchAllPolls, fetchSinglePoll } from "../services/PollService";
import { nextTick } from "vue";

const polls = ref([]);
const pollsDetails = ref({});

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
    console.log("Final poll details:", pollsDetails.value);
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
  <div class="flex flex-col items-center min-h-screen w-full bg-sky-800">
    <Header></Header>
  </div>
</template>
