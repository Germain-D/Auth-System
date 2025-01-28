// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "@nuxt/icon", "@pinia/nuxt",'pinia-plugin-persistedstate/nuxt'],
  runtimeConfig: {
    public: {
      GOOGLE_CLIENT_ID: process.env.GOOGLE_CLIENT_ID,
      GOOGLE_REDIRECT_URI: process.env.GOOGLE_REDIRECT_URI,
      LINKEDIN_CLIENT_ID: process.env.LINKEDIN_CLIENT_ID,
      LINKEDIN_REDIRECT_URI: process.env.LINKEDIN_REDIRECT_URI,
      GITHUB_CLIENT_ID: process.env.GITHUB_CLIENT_ID,
      GITHUB_REDIRECT_URI: process.env.GITHUB_REDIRECT_URI,
      FACEBOOK_CLIENT_ID: process.env.FACEBOOK_CLIENT_ID,
      FACEBOOK_REDIRECT_URI: process.env.FACEBOOK_REDIRECT_URI,
      BACKEND_URL: process.env.BACKEND_URL,
    }
  }
})