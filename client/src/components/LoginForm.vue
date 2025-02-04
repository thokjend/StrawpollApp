<script setup>
import { ref } from "vue";

const username = ref("");
const password = ref("");
const registerMode = ref(false);

const login = () => {
  console.log(username.value);
};
</script>
<template></template>
