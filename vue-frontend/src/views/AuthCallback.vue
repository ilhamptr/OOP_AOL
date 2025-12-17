<template>
  <div class="flex items-center justify-center h-screen">
    <p class="text-gray-600">Signing you in...</p>
  </div>
</template>

<script setup>
import { onMounted } from "vue"
import { useRouter } from "vue-router"

const router = useRouter()

onMounted(async () => {
  const hash = window.location.hash.substring(1)
  const params = new URLSearchParams(hash)

  const supabaseAccessToken = params.get("access_token")

  if (!supabaseAccessToken) {
    console.error("No access_token in URL")
    return router.push("/login")
  }

  try {
    const response = await fetch("http://127.0.0.1:8080/auth", {
      method: "POST",
      headers: {
        "Authorization": "Bearer " + supabaseAccessToken,
        "Content-Type": "application/json"
      }
    })

    const data = await response.json()

    if (!response.ok) {
      console.error("OAuth login failed", data)
      return router.push("/login")
    }

    // Store tokens
    localStorage.setItem("jwt", data.access_token)
    localStorage.setItem("refresh_token", data.refresh_token)

    // Redirect success
    router.push("/main")

  } catch (err) {
    console.error("OAuth error:", err)
    router.push("/login")
  }
})
</script>
