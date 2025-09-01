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
        <h2 class="text-xl font-semibold text-gray-800">Create an account</h2>
        <p class="text-gray-500 text-sm">Fill in your details to sign up</p>
      </div>

      <!-- Form -->
      <form @submit.prevent="handleSignup" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-700">Email</label>
          <input v-model="email" type="email" required placeholder="Enter your email"
            class="mt-1 w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none"/>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">Username</label>
          <input v-model="username" type="text" required placeholder="Choose a username"
            class="mt-1 w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none"/>
        </div>

        <div>
          <label class="block text-sm font-medium text-gray-700">Password</label>
          <input v-model="password" type="password" required placeholder="Enter a password"
            class="mt-1 w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-purple-500 focus:outline-none"/>
        </div>

        <button type="submit"
          class="w-full bg-purple-600 hover:bg-purple-700 text-white py-2 rounded-xl font-medium transition">
          Sign Up
        </button>
      </form>

      <!-- Footer -->
      <p class="mt-6 text-center text-sm text-gray-600">
        Already have an account?
        <RouterLink to="/login" class="text-purple-600 font-medium hover:underline">Sign in</RouterLink>
      </p>

      <!-- Status message -->
      <p v-if="message" class="mt-4 text-center text-sm" :class="messageColor">
        {{ message }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const email = ref("");
const username = ref("");
const password = ref("");
const message = ref("");
const messageColor = ref("text-gray-600");

const API_BASE = "http://127.0.0.1:8080";

// ✅ Handle signup
async function handleSignup() {
  try {
    const res = await fetch(`${API_BASE}/sign-up`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email: email.value,
        username: username.value,
        password: password.value
      }),
    });

    const data = await res.json();

    if (res.status === 201) {
      localStorage.setItem("jwt", data.token);

      message.value = "✅ Please check your email to verify your account.";
      messageColor.value = "text-green-600";

      // optional: redirect to login after delay
      setTimeout(() => router.push("/login"), 2000);
    } else if (res.status === 409) {
      message.value = "❌ Username or email already exists.";
      messageColor.value = "text-red-600";
    } else {
      message.value = data.message || "❌ Signup failed.";
      messageColor.value = "text-red-600";
    }
  } catch (err) {
    console.error(err);
    message.value = "❌ An error occurred. Please try again.";
    messageColor.value = "text-red-600";
  }
}
</script>
