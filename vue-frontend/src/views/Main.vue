<template>
  <div class="bg-gray-50 p-8 min-h-screen">
    <!-- Header -->
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold text-gray-800">My Job Postings</h1>
      <button
        @click="logout"
        class="px-4 py-2 bg-red-500 text-white rounded-lg shadow hover:bg-red-600"
      >
        Logout
      </button>
    </div>

    <!-- Stats -->
    <div class="p-6">
      <div class="mb-6 overflow-x-auto">
        <div class="flex gap-4 sm:grid sm:grid-cols-2 lg:grid-cols-4 min-w-max sm:min-w-0">
          <div class="bg-white p-6 rounded-lg shadow text-center flex-1 min-w-[200px]">
            <h3 class="text-lg font-semibold text-gray-700">Total Jobs</h3>
            <p class="text-2xl font-bold text-purple-600">{{ stats.total }}</p>
          </div>
          <div class="bg-white p-6 rounded-lg shadow text-center flex-1 min-w-[200px]">
            <h3 class="text-lg font-semibold text-gray-700">Active Jobs</h3>
            <p class="text-2xl font-bold text-green-600">{{ stats.active }}</p>
          </div>
          <div class="bg-white p-6 rounded-lg shadow text-center flex-1 min-w-[200px]">
            <h3 class="text-lg font-semibold text-gray-700">Inactive Jobs</h3>
            <p class="text-2xl font-bold text-red-600">{{ stats.inactive }}</p>
          </div>
          <div class="bg-white p-6 rounded-lg shadow text-center flex-1 min-w-[200px]">
            <h3 class="text-lg font-semibold text-gray-700">Latest Job</h3>
            <p class="text-sm font-medium text-gray-600">{{ stats.latest || '-' }}</p>
          </div>
        </div>
      </div>

      <!-- Job Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
            v-for="job in jobs"
            :key="job.ID"
            class="bg-white p-6 rounded-lg shadow"
            >
            <h2 class="text-lg font-semibold text-gray-800">{{ job.JobTitle }}</h2>
            <div class="flex gap-2 mt-3 flex-wrap">
                <span
                class="px-2 py-1 text-xs rounded"
                :class="job.Active ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'"
                >
                {{ job.Active ? 'Active' : 'Inactive' }}
                </span>
            </div>
            <p class="text-gray-600 mt-3">{{ job.Description.substring(0, 150) }}...</p>

            <div class="flex justify-between items-center mt-4 text-sm text-gray-500">
                <span>📅 {{ new Date(job.CreatedAt).toLocaleDateString() }}</span>
                <router-link
                :to="`/applicants/${job.ID}`"
                class="text-purple-600 font-medium"
                >
                View applicants
                </router-link>
            </div>

            <!-- Buttons -->
            <div class="flex gap-2 mt-3">
                <button
                @click="copyShareLink(job.ID)"
                class="px-3 py-1 text-sm bg-purple-600 text-white rounded hover:bg-purple-700"
                >
                📋 Copy Share Link
                </button>
                <button
                @click="deleteJob(job.ID)"
                class="px-3 py-1 text-sm bg-red-500 text-white rounded hover:bg-red-600"
                >
                🗑 Delete
                </button>
            </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center"
    >
      <div class="bg-white rounded-lg shadow-lg w-full max-w-lg p-6">
        <h2 class="text-xl font-bold mb-4">Post New Job</h2>
        <form @submit.prevent="addJob" class="space-y-4">
          <div>
            <label class="block text-gray-700 font-medium">Job Title</label>
            <input
              v-model="form.job_title"
              type="text"
              required
              class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring focus:border-purple-500"
            />
          </div>
          <div>
            <label class="block text-gray-700 font-medium">Description</label>
            <textarea
              v-model="form.description"
              rows="6"
              required
              class="w-full border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring focus:border-purple-500"
            ></textarea>
          </div>
          <div class="flex justify-end gap-3">
            <button
              type="button"
              @click="showModal = false"
              class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
            >
              Cancel
            </button>
            <button
              type="submit"
              class="px-4 py-2 bg-purple-600 text-white rounded hover:bg-purple-700"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Floating button -->
    <button
      @click="showModal = true"
      class="fixed bottom-6 right-6 bg-purple-600 text-white rounded-full shadow-lg w-14 h-14 flex items-center justify-center text-2xl hover:bg-purple-700"
      title="Post New Job"
    >
      +
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { apiRequest } from "../lib/api";

const router = useRouter();
const jobs = ref([]);
const stats = ref({ total: 0, active: 0, inactive: 0, latest: null });
const showModal = ref(false);
const form = ref({ job_title: "", description: "" });

function logout() {
  localStorage.removeItem("jwt");
  localStorage.removeItem("refresh_token");

  router.push("/login");
}


async function deleteJob(jobId) {
  const res = await apiRequest(`http://127.0.0.1:8080/delete-job/${jobId}`, {
    method: "DELETE",
  });

  if (!res.ok) {
    alert("❌ Failed to delete job");
    return;
  }

  jobs.value = jobs.value.filter((j) => j.ID !== jobId);
  alert("✅ Job deleted successfully!");
}



async function fetchJobs() {
  try {
    const res = await apiRequest("http://127.0.0.1:8080/jobs");

    if (!res.ok) {
      router.push("/login");
      return;
    }

    const data = await res.json();
    jobs.value = data;

    // Stats
    stats.value.total = data.length;
    stats.value.active = data.filter((j) => j.Active).length;
    stats.value.inactive = data.length - stats.value.active;

    if (data.length > 0) {
      const latestJob = data.reduce((a, b) =>
        new Date(a.CreatedAt) > new Date(b.CreatedAt) ? a : b
      );
      stats.value.latest = latestJob.JobTitle;
    }
  } catch (err) {
    console.error("Error fetching jobs:", err);
    router.push("/login");
  }
}


async function addJob() {
  try {
    const res = await apiRequest("http://127.0.0.1:8080/add-job", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(form.value),
    });

    if (!res.ok) {
      alert("Failed to create job");
      return;
    }

    showModal.value = false;
    form.value = { job_title: "", description: "" };

    await fetchJobs();
  } catch (error) {
    console.error("Error posting job:", error);
  }
}


function copyShareLink(jobId) {
  const link = `http://localhost:5173/apply/${jobId}`; // change this for prod
  navigator.clipboard.writeText(link).then(() => {
    alert("✅ Share link copied to clipboard!");
  });
}

onMounted(fetchJobs);
</script>
