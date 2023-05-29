<script setup lang="ts">

import { useSettingsStore } from '@/stores/settings'
import { decodeHexString, encodeHexString, getRemote } from '@/utils/url'
import { onMounted, onUnmounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const settings = useSettingsStore()
const files = ref<DirectoryEntry[]>([])

const route = useRoute()
const router = useRouter()

const rootFetcher = async () => {
  const res = await fetch(`${getRemote()}/api/walk`, {
    method: 'POST',
    body: JSON.stringify({
      'subdir': ''
    })
  })

  const data: APIResponse = await res.json()
  files.value = data.list
}

const fetcherSubfolder = async (path: string) => {
  const p = path.split('/')
  const subdir = p.length <= settings.basepathLength
    ? ''
    : `${p.slice(settings.basepathLength).join('/')}/`

  const res = await fetch(`${getRemote()}/api/walk`, {
    method: 'POST',
    body: JSON.stringify({
      'subdir': subdir
    })
  })

  let data: APIResponse = await res.json()

  const parent = p.length === settings.basepathLength
    ? ''
    : p.slice(settings.basepathLength).join('/')

  // ugly af
  if (parent) {
    data.list = [{
      name: '..',
      path: parent,
      isDirectory: true,
      isVideo: false,
      modTime: '',
      shaSum: '',
      size: 0,
      upperLevel: parent,
    }, ...data.list]
  }

  files.value = data.list
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
  // @ts-ignore
  () => fetcherSubfolder(decodeHexString(route.path.split('/').at(-1))),
)

onMounted(() => settings.fetchBasepathLength().then(rootFetcher))
onUnmounted(() => stop())
</script>

<template>
  <v-list lines="two">
    <v-list-item @click="onEntryClick(file)" v-for="file in files" :key="file.name" :title="file.name"
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