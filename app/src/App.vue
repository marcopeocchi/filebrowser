<script setup lang="ts">
import { ref } from 'vue';
import { RouterView } from 'vue-router'
import Drawer from './components/Drawer.vue';
import Toolbar from './components/Toolbar.vue';

const showDrawer = ref(false)
</script>

<template>
  <main>
    <Toolbar :on-drawer-click="() => showDrawer = !showDrawer" />
    <Drawer v-model="showDrawer" />

    <RouterView />
  </main>
</template>
