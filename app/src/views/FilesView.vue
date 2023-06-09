<script setup lang="ts">

import { useSettingsStore } from '@/stores/settings'
import { decodeHexString, encodeHexString, getRemote } from '@/utils/url'
import { onMounted, onUnmounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

import Loader from '@/components/FullPageLoader.vue'
import { useSearchStore } from '@/stores/search'

const settings = useSettingsStore()
const search = useSearchStore()

const loading = ref(false)
const files = ref<DirectoryEntry[]>([])
const filteredFiles = ref<DirectoryEntry[]>([])

const route = useRoute()
const router = useRouter()

const rootFetcher = async () => {
  loading.value = true
  const res = await fetch(`${getRemote()}/api/walk`, {
    method: 'POST',
    body: JSON.stringify({
      'subdir': ''
    })
  })

  if (!res.ok) {
    router.push('/login')
  }

  const data: APIResponse = await res.json()

  files.value = data.list
  filteredFiles.value = data.list
  loading.value = false
}

const fetcherSubfolder = async (path: string) => {
  const p = path.split('/')
  const subdir = p.length <= settings.basepathLength
    ? ''
    : `${p.slice(settings.basepathLength).join('/')}/`

  loading.value = true
  const res = await fetch(`${getRemote()}/api/walk`, {
    method: 'POST',
    body: JSON.stringify({
      'subdir': subdir
    })
  })

  if (!res.ok) {
    router.push('/login')
  }

  let data: APIResponse = await res.json()

  files.value = data.list
  filteredFiles.value = data.list
  loading.value = false
}

const openFile = (path: string) => {
  window.location.href = `${getRemote()}/api/open/${encodeHexString(path)}`
}

const onEntryClick = (entry: DirectoryEntry) => {
  if (entry.isDirectory) {
    fetcherSubfolder(entry.path)
    router.push(`/files/${encodeHexString(entry.path)}`)
    return
  }
  openFile(entry.path)
}

const getFormattedDate = (entry: DirectoryEntry) => entry.name === '..'
  ? ''
  : new Date(entry.modTime).toLocaleString()

const stop = watch(
  () => route.path,
  () => fetcherSubfolder(decodeHexString(route.path.split('/').pop()!)),
)

watch(
  search,
  (current) => {
    filteredFiles.value = files.value.filter(f => f.name
      .toLowerCase()
      .includes(current.search.toLowerCase())
    )
  }
)

onMounted(() => settings.fetchBasepathLength().then(rootFetcher))
onUnmounted(() => stop())
</script>

<template>
  <v-list lines="two">
    <Loader :open="loading" />
    <v-list-item @click="onEntryClick(file)" v-for="file in filteredFiles" :key="file.name" :title="file.name"
      :subtitle="getFormattedDate(file)">
      <template v-slot:prepend>
        <v-avatar color="grey-lighten-1">
          <v-icon color="white">
            {{ file.isDirectory ? 'mdi-folder' : 'mdi-file' }}
          </v-icon>
        </v-avatar>
      </template>
    </v-list-item>
  </v-list>
</template>