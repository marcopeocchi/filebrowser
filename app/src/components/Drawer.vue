<script setup lang="ts">
import ThemeToggler from '@/components/ThemeToggler.vue'
import { computed } from 'vue';

const props = defineProps(['modelValue'])
const emits = defineEmits(['update:modelValue'])

const value = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emits('update:modelValue', value)
  }
})
</script>

<template>
  <v-layout>
    <v-navigation-drawer v-model="value" location="bottom" temporary>
      <v-list-item>
        <template v-slot:prepend>
          <ThemeToggler />
        </template>
      </v-list-item>
    </v-navigation-drawer>
  </v-layout>
</template>