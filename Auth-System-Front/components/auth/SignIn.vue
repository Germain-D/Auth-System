<template>
  <AuthLoginForm />

  <hr class="my-4"></hr>

  <AuthButton @click="loginWithGoogle" size="lg" block styleName="google" class="mt-1">Se Connecter avec Google</AuthButton>
  <AuthButton @click="loginWithLinkedIn" size="lg" block styleName="linkedin" class="mt-1">Se Connecter avec LinkedIn</AuthButton>
<!-- <AuthButton size="lg" block styleName="github" class="mt-1">Se Connecter avec Github</AuthButton>
  <AuthButton size="lg" block styleName="facebook" class="mt-1">Se Connecter avec Facebook</AuthButton>

  <AuthButton size="lg" block styleName="apple" class="mt-1  "> Se Connecter avec Apple</AuthButton>--> 
</template>

<script setup>
import { useAuthStore } from '@/stores/auth';
const config = useRuntimeConfig()
const authStore = useAuthStore();

const loginWithGoogle = () => {
  // Google OAuth URL with your client ID
  const googleAuthUrl = 'https://accounts.google.com/o/oauth2/v2/auth';
  const clientId = config.public.GOOGLE_CLIENT_ID;
  const redirectUri = 'http://localhost:8000/auth/google/callback'; //adresse de mon back-end

  // Generate a random state value
  const state = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);

  const params = {
    client_id: clientId,
    redirect_uri: redirectUri,
    response_type: 'code',
    scope: 'email profile',
    access_type: 'offline',
    prompt: 'consent',
    state: state, // Add the state parameter
  };

  const queryString = new URLSearchParams(params).toString();
  window.location.href = `${googleAuthUrl}?${queryString}`;
};

const loginWithLinkedIn = () => {
  const linkedinAuthUrl = 'https://www.linkedin.com/oauth/v2/authorization';
  const clientId = config.public.LINKEDIN_CLIENT_ID;
  const redirectUri = 'http://localhost:8000/auth/linkedin/callback'; // URL de votre backend

  const state = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);

  const params = {
    client_id: clientId,
    redirect_uri: redirectUri,
    response_type: 'code',
    scope: 'r_liteprofile r_emailaddress', // Scopes pour LinkedIn
    state: state,
  };

  const queryString = new URLSearchParams(params).toString();
  window.location.href = `${linkedinAuthUrl}?${queryString}`;
};
</script>