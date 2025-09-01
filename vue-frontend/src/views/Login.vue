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
        <h1 class="text-2xl font-bold text-gray-900 mt-4">SecureApp</h1>
        <p class="text-gray-500 text-sm">Your trusted platform for secure access</p>
      </div>

      <!-- Welcome -->
      <div class="text-center mb-6">
        <h2 class="text-xl font-semibold text-gray-800">Welcome back</h2>
        <p class="text-gray-500 text-sm">Enter your credentials to access your account</p>
      </div>

      <!-- Form -->
      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700">Email</label>
          <input v-model="email" type="email" required placeholder="Enter your email"
            class="mt-1 w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none"/>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">Username</label>
          <input v-model="username" type="text" required placeholder="Enter your username"
            class="mt-1 w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none"/>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">Password</label>
          <input v-model="password" type="password" required placeholder="Enter your password"
            class="mt-1 w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none"/>
        </div>

        <div class="flex items-center justify-between text-sm">
          <label class="flex items-center space-x-2">
            <input type="checkbox" v-model="remember" class="rounded"/>
            <span class="text-gray-600">Remember me</span>
          </label>
          <RouterLink to="/forgot-password" class="text-purple-600 hover:underline">Forgot password?</RouterLink>
        </div>

        <button type="submit"
          class="w-full bg-purple-600 hover:bg-purple-700 text-white py-2 rounded-xl font-medium transition">
          Sign In
        </button>
      </form>

      <!-- Footer -->
      <p class="mt-6 text-center text-sm text-gray-600">
        Don’t have an account?
        <RouterLink to="/signup" class="text-purple-600 font-medium hover:underline">Sign up</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue"
import { useRouter } from "vue-router"

const email = ref("")
const username = ref("")
const password = ref("")
const remember = ref(false)

const router = useRouter()

const handleLogin = async () => {
  try {
    const response = await fetch("http://127.0.0.1:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email: email.value,
        username: username.value,
        password: password.value,
      }),
    })

    const data = await response.json()

    if (response.ok) {
      localStorage.setItem("jwt", data.token)

      // ✅ Redirect using Vue Router instead of window.location.href
      router.push("/main")
    } else {
      alert(data.message || "Login failed")
    }
  } catch (err) {
    console.error(err)
    alert("An error occurred. Please try again.")
  }
}
</script>
