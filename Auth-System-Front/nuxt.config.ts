// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "@nuxt/icon", "@pinia/nuxt",'pinia-plugin-persistedstate/nuxt'],
  runtimeConfig: {
    public: {
      GOOGLE_CLIENT_ID: process.env.GOOGLE_CLIENT_ID,
      LINKEDIN_CLIENT_ID: process.env.LINKEDIN_CLIENT_ID,
      GITHUB_CLIENT_ID: process.env.GITHUB_CLIENT_ID,
    }
  }
})