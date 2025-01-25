import { defineStore } from 'pinia';

interface AuthState {
  isAuthenticated: boolean;
  error: string | null;
}

interface Credentials {
  email: string;
  password: string;
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    isAuthenticated: false,
    error: null
  }),

  actions: {
    async login(credentials: Credentials) {
      try {
        const response = await fetch('http://localhost:8000/api/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include', // Important for cookies
          body: JSON.stringify(credentials),
        });

        const data = await response.json();
        
        if (!response.ok) {
          this.error = data.error || 'Login failed';
          this.isAuthenticated = false;
          throw new Error(this.error || 'Unknown error');
        }

        this.isAuthenticated = true;
        this.error = null;
        return true;
      } catch (error) {
        this.error = error instanceof Error ? error.message : 'Unknown error';
        this.isAuthenticated = false;
        return false;
      }
    },

    setToken(token: string | null) {
      this.isAuthenticated = !!token;
      if (token) {
        localStorage.setItem('jwt', token);
      } else {
        localStorage.removeItem('jwt');
      }
    },

    async logout() {
      try {
        await fetch('http://localhost:8000/api/logout', {
          method: 'POST',
          credentials: 'include',
        });
      } finally {
        this.setToken(null);
        this.isAuthenticated = false;
      }
    },

    async checkAuth() {
      try {
        const response = await fetch('http://localhost:8000/api/user', {
          credentials: 'include',
        });
        console.log(response);
        if (response.ok) {
          const data = await response.json();
          this.isAuthenticated = true;
          return true;
        }
      } catch (error) {
        this.setToken(null);
      }
      return false;
    }
  },

  getters: {

    getError: (state) => state.error,
    isLoggedIn: (state) => state.isAuthenticated
  }
});