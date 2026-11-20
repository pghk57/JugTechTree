<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchHealth, type Health } from '@/api/health'

const health = ref<Health>()
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    health.value = await fetchHealth()
  } catch (e) {
    error.value = String(e)
  }
})
</script>

<template>
  <h1>You did it!</h1>
  <p>
    Visit <a href="https://vuejs.org/" target="_blank" rel="noopener">vuejs.org</a> to read the
    documentation
  </p>

  <p>health check : {{ health == undefined ? error : health.res }}</p>
</template>

<style scoped></style>
