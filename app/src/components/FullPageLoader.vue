<script setup lang="ts">
type Props = {
  open: boolean
}

defineProps<Props>()
</script>

<template>
  <div class="text-center loader-center" v-if="open">
    <v-progress-circular indeterminate color="primary" />
  </div>
</template>

<style scoped>
.loader-center {
  height: 90vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>