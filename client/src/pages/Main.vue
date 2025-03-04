<script setup>
import { onMounted, ref } from "vue";
import { ValidateToken } from "../utils/ValidateToken";
import Header from "../components/Header.vue";
import { fetchAllPolls, fetchSinglePoll } from "../services/PollService";

const polls = ref([]);
const pollsDetails = ref([]);

onMounted(async () => {
  await fetchPollId();
  await getPolls();
  console.log(pollsDetails.value);
});

const getPolls = async () => {
  try {
    // Fetch details for each poll ID
    const pollData = await Promise.all(
      polls.value.map(async (id) => {
        const response = await fetchSinglePoll(id);
        return response.data; // Extract `data` from the API response
      })
    );
    pollDetails.value = pollData; // Store all fetched poll details
  } catch (error) {
    console.log("Error fetching poll details");
  }
};

const fetchPollId = async () => {
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
