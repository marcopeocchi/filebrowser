import { getRemote } from '@/utils/url'
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  const basepathLength = ref(0)

  function setBasepathLength(value: number) {
    basepathLength.value = value
  }

  async function fetchBasepathLength() {
    const res = await fetch(`${getRemote()}/api/basepath/length`)
    const data: number = await res.json()

    basepathLength.value = data + 1
  }

  return { basepathLength, setBasepathLength, fetchBasepathLength }
})
