<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-b from-white to-purple-50">
    <div class="w-full max-w-md bg-white shadow-xl rounded-2xl p-8">

      <!-- Logo -->
      <div class="flex flex-col items-center mb-6">
        <div class="bg-purple-600 p-4 rounded-2xl">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M12 11c.5304 0 1.0391-.2107 1.4142-.5858C13.7893 10.0391 14 9.5304 14 9V7c0-1.1046-.8954-2-2-2s-2 .8954-2 2v2c0 .5304.2107 1.0391.5858 1.4142C10.9609 10.7893 11.4696 11 12 11zM6 9v12h12V9H6z"/>
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-gray-900 mt-4">MatchGrid</h1>
        <p class="text-gray-500 text-sm">Your trusted platform for Recruitment Process</p>
      </div>

      <!-- Welcome -->
      <div class="text-center mb-6">
        <h2 class="text-xl font-semibold text-gray-800">Welcome back</h2>
        <p class="text-gray-500 text-sm">Enter your credentials to access your account</p>
      </div>

      <!-- Social Login -->
      <div class="space-y-3 mb-6">
        <button @click="loginWithGoogle"
          class="w-full border py-2 rounded-xl flex items-center justify-center gap-2 hover:bg-gray-100">
          <img src="https://www.svgrepo.com/show/355037/google.svg" class="h-5" />
          Continue with Google
        </button>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue"
import { useRouter } from "vue-router"
import { supabase } from "../lib/supabase"


const router = useRouter()

const loginWithGoogle = async () => {
  const { error } = await supabase.auth.signInWithOAuth({
    provider: "google",
    options: { redirectTo: window.location.origin + "/auth/callback" }
  })
  if (error) console.error(error)
}


</script>
