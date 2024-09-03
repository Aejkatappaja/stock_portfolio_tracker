<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <div class="min-h-screen bg-gray-800 p-4 text-white">
    <div>
      <RouterLink
        to="/"
        class="px-1 text-lg font-semibold transition duration-300 hover:text-gray-300"
        >Dashboard
      </RouterLink>
    </div>
    <div class="mt-8"><RouterView /></div>
  </div>
</template>

<style scoped></style>
