<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>

<template>
  <div class="min-h-screen bg-[#292828] p-4 text-white">
    <div>
      <RouterLink
        to="/"
        class="px-1 text-lg font-semibold text-[#a9b665] transition duration-300 hover:text-gray-300"
        >Dashboard
      </RouterLink>
    </div>
    <div class="mt-8"><RouterView /></div>
  </div>
</template>

<style scoped></style>
