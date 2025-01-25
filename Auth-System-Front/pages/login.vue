<script setup lang="ts">

definePageMeta({
  layout: 'landing',
})
</script>

<template>

    <div class="grid mx-auto max-w-sm gap-6">
      <div class="grid gap-2 text-center">
        <h1 class="text-2xl font-semibold tracking-tight">
          Se Connecter
        </h1>
        <p class="text-balance text-sm text-muted-foreground">
            Connectez-vous avec le système d'authentification de votre choix
        </p>
      </div>
      <AuthSignIn />
    </div>

</template>