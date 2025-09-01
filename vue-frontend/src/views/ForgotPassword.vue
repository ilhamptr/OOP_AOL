<template>
  <div class="bg-purple-50 flex items-center justify-center min-h-screen">
    <div class="bg-white shadow-lg rounded-lg p-8 w-full max-w-md">
      
      <!-- Step 1: Forgot Password -->
      <div v-if="step === 'forgot'">
        <h2 class="text-2xl font-bold text-purple-700 mb-4">Forgot Password</h2>
        <p class="text-gray-600 text-sm mb-6">
          Enter your email address and we’ll send you an OTP to reset your password.
        </p>

        <form @submit.prevent="sendOtp" class="space-y-4">
          <input 
            v-model="email"
            type="email" 
            placeholder="Enter your email" 
            class="w-full border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-500" 
            required 
          />
          <button 
            type="submit" 
            class="w-full bg-purple-600 text-white py-2 rounded hover:bg-purple-700"
          >
            Send OTP
          </button>
        </form>
      </div>

      <!-- Step 2: OTP + New Password -->
      <div v-else>
        <h2 class="text-2xl font-bold text-purple-700 mb-4">Reset Password</h2>
        
        <p class="text-gray-600 text-sm mb-2">Enter the 6-digit OTP sent to your email</p>
        <div class="flex justify-between mb-4">
          <input 
            v-for="(digit, idx) in otpInputs"
            :key="idx"
            v-model="otpInputs[idx]"
            maxlength="1"
            type="text"
            class="w-12 h-12 border rounded text-center text-lg"
            @input="focusNext(idx)"
            @keydown.backspace="focusPrev(idx, $event)"
            required
          />
        </div>

        <p class="text-gray-600 text-sm mb-2">Set your new password</p>
        <form @submit.prevent="verifyOtpAndReset" class="space-y-3">
          <input 
            v-model="newPassword"
            type="password" 
            placeholder="New Password" 
            class="w-full border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-500" 
            required 
          />
          <input 
            v-model="confirmPassword"
            type="password" 
            placeholder="Confirm New Password" 
            class="w-full border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-purple-500" 
            required 
          />
          <button 
            type="submit" 
            class="w-full bg-purple-600 text-white py-2 rounded hover:bg-purple-700"
          >
            Verify & Reset Password
          </button>
        </form>

        <p class="text-sm text-gray-500 text-center mt-3">
          Didn’t get the OTP? 
          <span @click="resendOtp" class="text-purple-600 hover:underline cursor-pointer">Send again</span>
        </p>
      </div>

      <!-- Status Message -->
      <div v-if="message" class="mt-4 text-sm" :class="messageColor">
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const email = ref("");
const step = ref("forgot"); // "forgot" | "reset"
const otpInputs = ref(["", "", "", "", "", ""]);
const newPassword = ref("");
const confirmPassword = ref("");
const message = ref("");
const messageColor = ref("text-gray-600");

const API_BASE = "http://127.0.0.1:8080";

// ✅ Send OTP
async function sendOtp() {
  try {
    const res = await fetch(`${API_BASE}/forgot-password`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email: email.value }),
    });
    if (!res.ok) throw new Error("Failed to send OTP");
    step.value = "reset";
    message.value = "✅ OTP sent to your email!";
    messageColor.value = "text-green-600";
  } catch (err) {
    message.value = "❌ Failed to send OTP.";
    messageColor.value = "text-red-600";
  }
}

// ✅ Resend OTP
async function resendOtp() {
  try {
    const res = await fetch(`${API_BASE}/forgot-password`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email: email.value }),
    });
    if (!res.ok) throw new Error("Resend failed");
    message.value = "🔄 OTP resent to your email!";
    messageColor.value = "text-green-600";
  } catch {
    message.value = "❌ Failed to resend OTP.";
    messageColor.value = "text-red-600";
  }
}

// ✅ Verify OTP & Reset Password
async function verifyOtpAndReset() {
  const otp = otpInputs.value.join("");

  // validation
  const passwordRegex = /^(?=.*[!@#$%^&*(),.?":{}|<>]).{8,}$/;
  if (!passwordRegex.test(newPassword.value)) {
    message.value = "❌ Password must be at least 8 chars & include 1 special character.";
    messageColor.value = "text-red-600";
    return;
  }
  if (newPassword.value !== confirmPassword.value) {
    message.value = "❌ Passwords do not match.";
    messageColor.value = "text-red-600";
    return;
  }

  try {
    const res = await fetch(`${API_BASE}/verify-otp`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ 
        email: email.value, 
        otp, 
        new_password: newPassword.value 
      }),
    });
    if (!res.ok) throw new Error("Reset failed");

    message.value = "🎉 Password reset successfully! You can now log in.";
    messageColor.value = "text-green-600";
    otpInputs.value = ["", "", "", "", "", ""];
    newPassword.value = "";
    confirmPassword.value = "";
  } catch {
    message.value = "❌ Invalid OTP or reset failed.";
    messageColor.value = "text-red-600";
  }
}

// ✅ OTP input auto focus
function focusNext(idx) {
  if (otpInputs.value[idx].length === 1 && idx < otpInputs.value.length - 1) {
    const nextInput = document.querySelectorAll("input")[idx + 1];
    nextInput?.focus();
  }
}
function focusPrev(idx, event) {
  if (event.key === "Backspace" && !otpInputs.value[idx] && idx > 0) {
    const prevInput = document.querySelectorAll("input")[idx - 1];
    prevInput?.focus();
  }
}
</script>
