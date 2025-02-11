<script setup>
const props = defineProps({
  type: {
    type: String,
    required: true,
  },
  options: {
    type: String,
    required: true,
  },
  multipleChoice: {
    type: String,
    required: true,
  },
});

const headerText = () => {
  return props.multipleChoice === "0"
    ? "Make a choice:"
    : "Choose as many as you like:";
};

const typeOfInput = () => {
  return props.multipleChoice === "0" ? "radio" : "checkbox";
};
</script>

<template>
  <div class="mt-4">
    <h2 class="font-semibold">{{ headerText() }}</h2>
    <ul class="mt-2">
      <li
        v-for="(choice, index) in options.split('|')"
        :key="index"
        class="p-2"
      >
        <div>
          <input
            class="cursor-pointer"
            :type="typeOfInput()"
            name="input"
            :id="'checkbox-' + index"
          /><label class="ml-2 cursor-pointer" :for="'checkbox-' + index">
            {{ choice }}
          </label>
        </div>
      </li>
    </ul>
  </div>
</template>
