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
      <AuthButton type="submit" size="lg" block>Se Connecter</AuthButton>
      <div id="result" class="mt-3 text-center"></div>
      <AuthButton type="submit" size="lg" block styleName="outline" class="my-3 hover:text-black" >Se Connecter</AuthButton>
      <AuthButton type="submit" size="lg" block styleName="google" class="my-3">Se Connecter avec Google</AuthButton>
      <AuthButton type="submit" size="lg" block styleName="github" class="my-3">Se Connecter avec Github</AuthButton>
      <AuthButton type="submit" size="lg" block styleName="facebook" class="my-3">Se Connecter avec Facebook</AuthButton>
      <AuthButton type="submit" size="lg" block styleName="linkedin" class="my-3">Se Connecter avec LinkedIn</AuthButton>
      <AuthButton type="submit" size="lg" block styleName="apple" class="my-3  "><span class="hover:text-black"> Se Connecter avec Apple</span></AuthButton>
      <AuthButton type="submit" size="lg" block styleName="circle" class="my-3">Se Connecter avec Apple</AuthButton>
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
      status.value = "loading";
      const response = await $fetch("https://api.web3forms.com/submit", {
        method: "POST",
        body: form.value,
      });
      console.log(response);
      result.value = response.message;
      if (response.status === 200) {
        status.value = "success";
      } else {
        console.log(response); // Log for debugging, can be removed
        status.value = "error";
      }
    } catch (error) {
      console.log(error); // Log for debugging, can be removed
      status.value = "error";
      result.value = "Something went wrong!";
    } finally {
      // Reset form after submission
      form.value.email = "";
        form.value.password = "";
  
      // Clear result and status after 5 seconds
      setTimeout(() => {
        result.value = "";
        status.value = "";
      }, 5000);
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