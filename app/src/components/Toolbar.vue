<script setup lang="ts">
type Props = {
  onDrawerClick: () => void
}

defineProps<Props>()
</script>

<template>
  <v-toolbar color="primary" class="rounded-0">
    <v-btn variant="text" icon="mdi-menu" @click.stop="onDrawerClick"></v-btn>

    <v-toolbar-title>File Browser</v-toolbar-title>
    <v-spacer></v-spacer>

    <v-btn variant="text" icon="mdi-magnify"></v-btn>
    <!-- <v-btn variant="text" icon="mdi-view-module"></v-btn> -->
  </v-toolbar>
</template>