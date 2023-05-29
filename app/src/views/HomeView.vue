<script setup lang="ts">
import Drawer from '@/components/Drawer.vue'
import Toolbar from '@/components/Toolbar.vue'
import { useSettingsStore } from '@/stores/settings'
import { getRemote } from '@/utils/url'
import { Buffer } from 'buffer'
import { onMounted, ref } from 'vue'

const settings = useSettingsStore()
const files = ref<DirectoryEntry[]>([])

const showDrawer = ref(false)

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
  const encoded = Buffer.from(path).toString('hex')
  window.open(`${getRemote()}/api/open/${encoded}`)
}

const onEntryClick = (entry: DirectoryEntry) => entry.isDirectory
  ? fetcherSubfolder(entry.path)
  : openFile(entry.path)

const getFormattedDate = (entry: DirectoryEntry) => entry.name === '..'
  ? ''
  : new Date(entry.modTime).toLocaleString()

onMounted(() => settings.fetchBasepathLength().then(rootFetcher))
</script>

<template>
  <main>
    <Toolbar :on-drawer-click="() => showDrawer = !showDrawer" />
    <Drawer v-model="showDrawer" />

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
        <!-- <template v-slot:append>
          <v-btn color="grey-lighten-1" icon="mdi-information" variant="text"></v-btn>
        </template> -->
      </v-list-item>
    </v-list>
  </main>
</template>