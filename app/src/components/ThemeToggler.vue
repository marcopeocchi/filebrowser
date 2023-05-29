<script setup lang="ts">
import { useTheme } from 'vuetify';

const theme = useTheme()

function toggleDarkMode() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
  localStorage.setItem('theme', theme.global.name.value)
}
</script>

<template>
  <v-switch label="Toggle theme" v-model="theme.global.current.value.dark" @click="toggleDarkMode"></v-switch>
</template>