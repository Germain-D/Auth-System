<template>
    <form @submit.prevent="submitForm">

      <div class="mb-5">
        <label for="email_address" class="sr-only">Adresse Email</label
        ><input
          id="email_address"
          type="email"
          placeholder="Adresse Email"
          name="email"
          required
          v-model="form.email"
          class="w-full px-4 py-3 border-2 placeholder:text-gray-800 rounded-md outline-none focus:ring-4 border-gray-300 focus:border-gray-600 ring-gray-100"
        />
        <div class="empty-feedback text-red-400 text-sm mt-1">
          Veuillez indiquer votre adresse email.
        </div>
        <div class="invalid-feedback text-red-400 text-sm mt-1">
          Veuillez fournir une adresse email valide.
        </div>
      </div>
      <div class="mb-5">
        <input
          type="text"
          placeholder="Password"
          required
          class="w-full px-4 py-3 border-2 placeholder:text-gray-800 rounded-md outline-none focus:ring-4 border-gray-300 focus:border-gray-600 ring-gray-100"
          name="password"
          v-model="form.password"
        />
        <div class="empty-feedback invalid-feedback text-red-400 text-sm mt-1">
            Veuillez indiquer votre mot de passe.
        </div>
      </div>
      <AuthButton type="submit" size="lg" block styleName="outline" class="my-3" >Se Connecter</AuthButton>
      <div id="result" class="mt-3 text-center"></div>
      

    </form>
  </template>
  
  <script setup>
  
  const form = ref({
    email: "",
    password: "",
  });
  
  const result = ref("");
  const status = ref("");
  
  const submitForm = async () => {
    try {
      const response = await fetch("http://localhost:8000/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(form.value),
      });
      const data = await response.json();
      if (response.ok) {
        status.value = "success";
        result.value = data.message;
        console.log(data);
        window.location.href = "/";
      } else {
        status.value = "error";
        result.value = data.error;
        console.log(data);
      }
    } catch (error) {
      status.value = "error";
      result.value = error.message;
      console.error(error);
    }
  };
  </script>
  
  <style>
  .invalid-feedback,
  .empty-feedback {
    display: none;
  }
  
  .was-validated :placeholder-shown:invalid ~ .empty-feedback {
    display: block;
  }
  
  .was-validated :not(:placeholder-shown):invalid ~ .invalid-feedback {
    display: block;
  }
  
  .is-invalid,
  .was-validated :invalid {
    border-color: #dc3545;
  }
  </style>