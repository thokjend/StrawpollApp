<script setup>
import { useRoute } from "vue-router";
import { onMounted, ref } from "vue";
import { PieChart } from "echarts/charts";
import VChart from "vue-echarts";
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import router from "../router";
import { ValidateToken } from "../utils/ValidateToken";

const route = useRoute();
const pollId = route.params.id;
const pollData = ref(null);
const ws = ref(null);

/* const colors = [
  "#FF0000", // Red
  "#00FF00", // Green
  "#0000FF", // Blue
  "#FFFF00", // Yellow
  "#FF00FF", // Magenta
]; */

onMounted(async () => {
  await ValidateToken();
  await fetchPollData();

  ws.value = new WebSocket("ws://localhost:8080/ws");

  ws.value.onopen = () => {
    console.log("WebSocket connected");
  };

  ws.value.onmessage = async (event) => {
    console.log("WebSocket message received:", event.data);
    if (event.data === "update") {
      await fetchPollData(); // Re-fetch poll data when an update is received
    }
  };

  ws.value.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  ws.value.onclose = () => {
    console.log("WebSocket disconnected");
  };
});

const fetchPollData = async () => {
  try {
    const response = await fetch(`http://localhost:8080/poll/${pollId}`);
    const result = await response.json();
    pollData.value = result.data;

    const votes = pollData.value.votes;

    const pieData = Object.keys(votes)
      .filter((key) => votes[key] > 0)
      .map((key) => ({
        name: key,
        value: parseInt(votes[key]),
        /* itemStyle: { color: colors[index % colors.length] }, */
      }));

    chartOptions.value.series[0].data = pieData;
    //console.log(pollData.value);
  } catch (error) {
    console.log("Error fetching poll");
  }
};

use([PieChart, CanvasRenderer]);
const chartOptions = ref({
  series: [
    {
      type: "pie",
      label: {
        show: true,
        position: "outside",
        formatter: "{b}: {c} ({d}%)",
      },
      data: [],
    },
  ],
});

const GetAllVotes = () => {
  let totalVotes = 0;
  const votes = Object.values(pollData?.value.votes || {});
  for (const vote of votes) {
    totalVotes += parseInt(vote);
  }
  return totalVotes;
};
</script>

<template>
  <div
    class="flex flex-col items-center justify-center min-h-screen w-full bg-sky-800 p-6"
  >
    <div
      class="border p-6 rounded-lg shadow-lg bg-white bg-opacity-80 max-w-4xl w-full flex flex-col md:flex-row space-y-6 md:space-y-0 md:space-x-6"
    >
      <div class="flex-1 flex flex-col">
        <h1 class="font-bold text-2xl mb-6 text-center md:text-left">
          {{ pollData?.title }}
        </h1>
        <div class="flex flex-col space-y-4">
          <div
            class="border p-4 rounded-md bg-gray-100"
            v-for="(vote, index) in Object.entries(pollData?.votes || {})"
            :key="index"
          >
            <div class="text-lg font-medium flex justify-between">
              <div>{{ vote[0] }}</div>
              <div>
                {{
                  GetAllVotes() > 0 && vote[1] > 0
                    ? ((vote[1] / GetAllVotes()) * 100).toFixed(2)
                    : "0"
                }}% ({{ vote[1] }} votes)
              </div>
            </div>

            <div class="w-full bg-gray-300 rounded-full h-5 mt-2">
              <div
                class="bg-blue-500 h-5 rounded-full transition-all duration-300"
                :style="{
                  width:
                    GetAllVotes() > 0
                      ? (vote[1] / GetAllVotes()) * 100 + '%'
                      : 0,
                }"
              ></div>
            </div>
          </div>
        </div>
        <div class="mt-auto flex p-4">
          <button
            @click="router.push(`/poll/${pollId}`)"
            class="bg-blue-600 text-white rounded w-25 h-10 font-semibold hover:bg-blue-700 transition duration-300 cursor-pointer"
          >
            Back to poll
          </button>
        </div>
      </div>
      <div class="flex-1 flex items-center justify-center">
        <div style="width: 100%; height: 400px">
          <v-chart class="chart" :option="chartOptions" />
        </div>
      </div>
    </div>
  </div>
</template>
