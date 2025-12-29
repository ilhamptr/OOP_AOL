<template>
  <div class="min-h-screen w-full flex flex-col lg:flex-row">
    <!-- Left Side: Content -->
    <div class="w-full lg:w-1/2 bg-[#EBF4F6] flex flex-col items-center justify-center p-8">
      <div class="flex flex-col items-center gap-10">
        <img src="../assets/matchgrid.svg" class="w-50 h-10" alt="Logo">

        <div class="text-center">
          <h1 class="text-4xl font-bold text-black mb-2">Welcome Back</h1>
          <p class="text-base text-[#AEAEAE]">Please sign in to access your account</p>
        </div>
        
        <div class="relative group cursor-pointer inline-flex rounded-[14px] overflow-hidden transition-transform bg-transparent p-0.5 drop-shadow-lg">
          <div 
            class="absolute inset-[-150%] aspect-square animate-[spin_3s_linear_infinite] 
                  bg-[conic-gradient(from_0deg,#EA4335,#FBBC04,#34A853,#4285F4,#EA4335)]"
          ></div>

          <button @click="loginWithGoogle" class="relative flex cursor-pointer items-center justify-center gap-3 px-20 py-3 bg-white rounded-xl transition-all group-hover:bg-gray-50 m-0 border-0">
            <!-- Google G Icon SVG -->
            <svg class="w-6 h-6" viewBox="0 0 24 24">
              <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#071952"/>
              <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#071952"/>
              <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#071952"/>
              <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#071952"/>
            </svg>
            <span class="text-lg font-semibold text-[#071952]">Login with Google</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Right Side: Black Box -->
    <div class="w-full lg:w-1/2 bg-[#EBF4F6] flex items-center justify-center">
      <img src="../assets/login-image.png"
     class="max-w-full max-h-[90vh] object-contain p-4 drop-shadow-lg">
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
