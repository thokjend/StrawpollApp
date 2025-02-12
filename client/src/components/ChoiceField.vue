<script setup>
import { ref } from "vue";
const props = defineProps({
  type: {
    type: String,
    required: true,
  },
  options: {
    type: Object,
    required: true,
  },
  multipleChoice: {
    type: String,
    required: true,
  },
});

const modelValue = defineModel();

const headerText = () => {
  return props.multipleChoice === "0"
    ? "Make a choice:"
    : "Choose as many as you like:";
};

const typeOfInput = () => {
  return props.multipleChoice === "0" ? "radio" : "checkbox";
};

// Function to handle selection logic
const handleSelection = (event, choice) => {
  if (props.multipleChoice === "1") {
    // Multiple choices (checkbox)
    if (event.target.checked) {
      modelValue.value.push(choice);
    } else {
      modelValue.value = modelValue.value.filter((item) => item !== choice);
    }
  } else {
    // Single choice (radio)
    modelValue.value = [choice];
  }
};
</script>

<template>
  <div class="mt-4">
    <h2 class="font-semibold">{{ headerText() }}</h2>
    <ul class="mt-2">
      <li
        v-for="(choice, index) in Object.keys(props.options)"
        :key="index"
        class="p-2"
      >
        <div>
          <input
            class="cursor-pointer"
            :type="typeOfInput()"
            name="input"
            :id="'checkbox-' + index"
            :value="choice"
            :checked="modelValue.includes(choice)"
            @change="handleSelection($event, choice)"
          /><label class="ml-2 cursor-pointer" :for="'checkbox-' + index">
            {{ choice }}
          </label>
        </div>
      </li>
    </ul>
  </div>
</template>
