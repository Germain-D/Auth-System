<script setup>
definePageMeta({
  middleware: ['guard']
})

import { useAuthStore } from '@/stores/auth'
const authStore = useAuthStore()

const handleLogout = async () => {
  await authStore.logout()
  navigateTo('/login')
}
</script>

<template>
  <div class="grid mx-auto max-w-sm gap-6 text-center py-8">
    <h1 class="text-2xl font-bold">Protected Page</h1>
    
    <div class="bg-green-100 border-l-4 border-green-500 text-green-700 p-4" role="alert">
      <p>Congratulations! You have successfully accessed this protected page.</p>
    </div>

    <button 
      @click="handleLogout"
      class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
    >
      Logout
    </button>
  </div>
</template>