<script setup lang="ts">
import { getRemote } from '@/utils/url'
import { useRouter } from 'vue-router'

type Props = {
  onDrawerClick: () => void
}

defineProps<Props>()

const router = useRouter()

const logout = async () => {
  const res = await fetch(`${getRemote()}/api/logout`)
  if (res.ok) {
    router.push('/login')
  }
}
</script>

<template>
  <v-toolbar color="primary" class="rounded-0">
    <v-btn variant="text" icon="mdi-menu" @click.stop="onDrawerClick" />

    <v-toolbar-title>File Browser</v-toolbar-title>
    <v-spacer />

    <v-btn variant="text" icon="mdi-magnify" />
    <v-btn variant="text" icon="mdi-logout" @click="logout" />
  </v-toolbar>
</template>