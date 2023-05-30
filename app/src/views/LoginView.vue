<script setup lang="ts">
import { getRemote } from '@/utils/url'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const form = ref(false)
const loading = ref(false)

const username = ref('')
const password = ref('')

const login = async () => {
  loading.value = true
  const res = await fetch(`${getRemote()}/api/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      username: username.value,
      password: password.value,
    })
  })
  loading.value = false
  if (res.ok) {
    router.push('/files')
  }
}
</script>

<template>
  <v-card class="mx-auto py-6 px-6 mt-8" max-width="600">
    <v-form v-model="form" @submit.prevent="">
      <v-text-field placeholder="Username" v-model="username" clearable :loading="loading" />
      <v-text-field placeholder="Password" v-model="password" clearable :loading="loading" />
      <v-btn class="mt-6" :disabled="!form" block color="success" size="large" type="submit" variant="elevated"
        @click="login">
        Sign In
      </v-btn>
    </v-form>
  </v-card>
</template>