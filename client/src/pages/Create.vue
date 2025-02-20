<script setup>
import { onMounted, ref } from "vue";
import router from "../router";

const options = ref([
  {
    name: "Option 1",
    value: "",
  },
  {
    name: "Option 2",
    value: "",
  },
]);
const toggleDescription = ref(false);
const title = ref("");
const description = ref("");
const multipleOptions = ref(false);
const reqNames = ref(false);
const missingTitle = ref(false);
const missingOptions = ref(false);
const sessionToken = ref("");

const deleteTask = () => {
  options.value.pop();
};

const addOption = () => {
  const correctNumber =
    parseInt(options.value[options.value.length - 1].name.split(" ")[1]) + 1;
  options.value.push({ name: `Option ${correctNumber}`, value: "" });
};

const createPoll = async () => {
  const filteredOptions = options.value
    .map((options) => options.value.trim())
    .filter((value) => value !== "");

  if (title.value.trim() === "") {
    missingTitle.value = true;
    return;
  } else {
    missingTitle.value = false;
  }

  if (filteredOptions.length === 0) {
    missingOptions.value = true;
    return;
  } else {
    missingOptions.value = false;
  }
  try {
    const response = await fetch("http://localhost:8080/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Title: title.value,
        Description: description.value,
        Options: filteredOptions,
        Settings: [multipleOptions.value, reqNames.value],
      }),
    });

    const result = await response.json();
    console.log("Poll created:", result);

    if (result.id) {
      router.push(`/poll/${result.id}`);
    }
  } catch (error) {
    console.error("Error", error);
  }
};

onMounted(async () => {
  const username = localStorage.getItem("username");
  try {
    const response = await fetch(`http://localhost:8080/session/${username}`);
    const result = await response.json();
    sessionToken.value = result;
  } catch (error) {
    console.log("Error fetching session token");
  }
});
</script>
<template>
  <div class="flex flex-col items-center min-h-screen w-full bg-sky-800 py-10">
    <h1 class="mb-4 text-3xl font-bold text-white">Create a poll</h1>
    <form
      @submit.prevent="createPoll"
      class="border p-6 rounded-lg shadow-lg bg-white bg-opacity-80 max-w-xl w-full"
    >
      <h1>Title</h1>
      <input
        class="border rounded p-2 w-full mt-2 focus:outline-none focus:ring-2 mb-2"
        :class="{
          'focus:ring-blue-500': !missingTitle,
          'focus:ring-red-500 border-red-500': missingTitle,
        }"
        type="text"
        placeholder="Type your question here"
        v-model="title"
      />

      <div v-if="missingTitle === true" class="text-red-500 font-bold">
        Please enter a poll title
      </div>

      <div v-if="toggleDescription === false" class="">
        <div class="flex">
          <div
            @click="toggleDescription = true"
            class="hover:underline cursor-pointer"
          >
            Add description
          </div>
        </div>

        <hr class="my-4 border-gray-950" />
      </div>

      <div v-if="toggleDescription === true">
        <h1 class="">Description (optional)</h1>
        <textarea
          class="w-full border rounded px-1 py-1"
          v-model="description"
        ></textarea>
        <div class="flex justify-end">
          <div
            @click="toggleDescription = false"
            class="hover:underline cursor-pointer"
          >
            Hide description
          </div>
        </div>
      </div>
      <h1>Answer Options</h1>
      <div class="relative w-full" v-for="option in options">
        <input
          class="border rounded p-2 w-full mt-2 focus:outline-none focus:ring-2 focus:ring-blue-500 mb-2"
          type="text"
          v-model="option.value"
          :placeholder="option.name"
        />
        <button
          @click.prevent="deleteTask"
          v-if="options.length > 1"
          class="absolute right-2 top-1/2 transform -translate-y-1/2 px-2 py-1 rounded cursor-pointer"
        >
          X
        </button>
      </div>
      <div>
        <button
          @click.prevent="addOption"
          class="bg-blue-600 text-white rounded w-full h-10 font-bold hover:bg-blue-700 transition duration-300 cursor-pointer"
        >
          Add option
        </button>
      </div>
      <div class="mt-4 p-4 border rounded-lg shadow-md">
        <h2 class="text-xl font-semibold mb-3">Settings</h2>
        <div class="space-y-3">
          <div class="flex items-center justify-between">
            <label for="multiple-options" class="text-gray-700 cursor-pointer"
              >Allow multiple options</label
            >
            <input
              id="multiple-options"
              class="w-5 h-5 accent-blue-600 cursor-pointer"
              type="checkbox"
              v-model="multipleOptions"
            />
          </div>
          <div class="flex items-center justify-between">
            <label for="require-names" class="text-gray-700 cursor-pointer"
              >Require names</label
            >
            <input
              id="require-names"
              class="w-5 h-5 accent-blue-600 cursor-pointer"
              type="checkbox"
              v-model="reqNames"
            />
          </div>
        </div>
      </div>
      <div
        v-if="missingOptions === true"
        class="mt-4 p-4 rounded-lg shadow-md bg-red-200 text-red-700"
      >
        Please enter at least one answer option
      </div>

      <button
        type="submit"
        class="bg-blue-600 text-white rounded w-full h-10 font-bold hover:bg-blue-700 transition duration-300 cursor-pointer mt-4"
      >
        Create poll
      </button>
    </form>
  </div>
</template>
