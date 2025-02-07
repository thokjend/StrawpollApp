<script setup>
const props = defineProps({
  type: {
    type: String,
    required: true,
  },
  id: {
    type: String,
    required: true,
  },
  label: {
    type: String,
    required: true,
  },
  modelValue: {
    type: String,
    required: true,
  },
  placeholder: {
    type: String,
    default: "",
  },
});
</script>

<template>
  <div class="flex items-center mb-4">
    <label :for="id" class="block text-lg font-bold pr-3 w-24 text-gray-700">
      {{ label }}:
    </label>
    <input
      :type="type"
      :id="id"
      class="border rounded p-2 w-full mt-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      :placeholder="placeholder"
    />
  </div>
</template>
