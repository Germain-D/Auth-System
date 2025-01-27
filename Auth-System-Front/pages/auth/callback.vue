<template>
    <div>
      <p>Connexion en cours...</p>
    </div>
  </template>
  
  <script setup>
  import { onMounted } from 'vue';
  import { useRouter } from 'vue-router';
  import { useAuthStore } from '@/stores/auth';
  
  const router = useRouter();
  const authStore = useAuthStore();
  
  onMounted(() => {
    const token = new URLSearchParams(window.location.search).get('token');
    
    if (token) {
      authStore.setToken(token);
      router.push('/only-authenticated');
    } else {
      console.error('Token missing in URL');
      router.push('/login');
    }
  });
  </script>