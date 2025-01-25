<script setup lang="ts">
definePageMeta({
  layout: 'landing',
})

import { useAuthStore } from "@/stores/auth";
import { computed, onMounted } from 'vue';

const authStore = useAuthStore();

// Make it reactive with computed
const connection_status = computed(() => authStore.isLoggedIn);

// Check auth status when component mounts
onMounted(async () => {
  await authStore.checkAuth();
});
</script>


<template>

    <div class="grid mx-auto max-w-sm gap-6">
      <h1 class="text-2xl font-bold text-center">Welcome to the Auth System</h1>
      <p class="text-center">This project serves as a template for a register/login system.</p>
      <p class="text-center">You can use this template to quickly set up authentication in your applications.</p>
      <h2 v-if="!connection_status" class="text-xl font-bold text-center">Vous n'êtes pas connecté</h2>
      <h2 v-if="connection_status" class="text-xl font-bold text-center">Vous êtes connecté</h2>
    </div>

</template>

