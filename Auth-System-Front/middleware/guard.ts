import { useAuthStore } from "@/stores/auth";

export default defineNuxtRouteMiddleware((to) => {
    const authStore = useAuthStore();
    
    // List of public routes that don't require authentication
    const publicRoutes = ['/login', '/register', '/'];
  
    console.log('Route middleware', to.path);
    console.log('Is logged in?', authStore.isLoggedIn);
    // If route requires auth and user isn't authenticated
    if (!publicRoutes.includes(to.path) && !authStore.isLoggedIn) {
      // Redirect to login
      return navigateTo('/login');
    }
  });