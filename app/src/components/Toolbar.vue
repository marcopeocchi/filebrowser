<script setup lang="ts">
import { useSearchStore } from '@/stores/search';
import { getRemote } from '@/utils/url'
import { ref } from 'vue';
import { useRouter } from 'vue-router'

type Props = {
  onDrawerClick: () => void
}

defineProps<Props>()

const emit = defineEmits(['search'])
const searchStore = useSearchStore()

const router = useRouter()

const searchExpand = ref(false)

const handleSearch = (event: any) => {
  emit('search', event.currentTarget.value)
  searchStore.setSearch(event.currentTarget.value)
}

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

    <v-btn variant="text" icon="mdi-magnify" @click="searchExpand = !searchExpand" />

    <v-expand-x-transition>
      <v-card flat class="mx-auto" color="transparent" width="400" v-show="searchExpand">
        <v-card-text>
          <v-text-field density="compact" variant="solo" label="Search files or folders" single-line hide-details
            @input="handleSearch">
          </v-text-field>
        </v-card-text>
      </v-card>
    </v-expand-x-transition>

    <v-btn variant="text" icon="mdi-logout" @click="logout" />
  </v-toolbar>
</template>