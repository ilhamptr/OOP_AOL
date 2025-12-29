<template>

  <div class="bg-white">
    <div class="p-4 border-b border-gray-200">
      <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div class="space-x-4">
          <span class="font-medium">Dashboard</span>
        </div>
        <div class="flex items-center justify-between sm:justify-end gap-4">
          <button 
            @click="showModal = true"
            class="bg-[linear-gradient(180deg,#071952_58%,#2F4A97_100%)] text-white px-6 py-2.5 rounded-xl font-medium hover:opacity-90 transition shadow-md"
          >
            Create Jobs
          </button>
          
          <div class="h-6 w-px bg-[#F2F2F2]"></div>

          <div class="relative">
            <div 
              @click="showProfileMenu = !showProfileMenu"
              class="w-11 h-11 bg-blue-100 rounded-xl flex items-center justify-center text-[#0A1B3F] cursor-pointer hover:bg-blue-200 transition select-none"
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-7 h-7">
                <path fill-rule="evenodd" d="M7.5 6a4.5 4.5 0 1 1 9 0 4.5 4.5 0 0 1-9 0ZM3.751 20.105a8.25 8.25 0 0 1 16.498 0 .75.75 0 0 1-.437.695A18.683 18.683 0 0 1 12 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 0 1-.437-.695Z" clip-rule="evenodd" />
              </svg>
            </div>

            <!-- Profile Dropdown -->
            <div 
              v-if="showProfileMenu" 
              class="absolute right-0 mt-2 w-48 bg-white rounded-xl shadow-lg border border-gray-100 py-1 z-50 origin-top-right"
            >

              <button 
                @click="logout" 
                class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 flex items-center gap-2 transition"
              >
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="w-4 h-4"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" x2="9" y1="12" y2="12"/></svg>
                Sign out
              </button>
            </div>
            
            <!-- Backdrop to close -->
            <div v-if="showProfileMenu" @click="showProfileMenu = false" class="fixed inset-0 z-40 cursor-default"></div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="min-h-screen bg-[#F2F4F9]">
    <div class="p-4 space-y-6">
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        <div>
          <!-- BACKEND: Stats Cards - Fetch summary counts from API -->
          <div class="bg-gradient-to-t from-white from-40% to-[#C9D8F6] to-100% p-4 rounded-xl space-y-6">
            <div class="flex justify-between">
              <span class="text-black text-2xl font-semibold">{{stats.total}}</span>
              <div class="bg-[#071952] rounded-lg p-2"></div>
            </div>
            <span class="text-[#8F8F8F] text-lg">Posted Jobs</span>
          </div>
        </div>
        <div>
          <div class="bg-white p-4 rounded-xl space-y-6">
            <div class="flex justify-between">
              <span class="text-black text-2xl font-semibold">{{stats.active}}</span>
              <div class="bg-[#071952] rounded-lg p-2"></div>
            </div>
            <span class="text-[#8F8F8F] text-lg">Active Jobs</span>
          </div>
        </div>
        <div>
          <div class="bg-white p-4 rounded-xl space-y-6">
            <div class="flex justify-between">
              <span class="text-black text-2xl font-semibold">{{ stats.inactive }}</span>
              <div class="bg-[#071952] rounded-lg p-2"></div>
            </div>
            <span class="text-[#8F8F8F] text-lg">Inactive Jobs</span>
          </div>
        </div>
        <div>
          <div class="bg-white p-4 rounded-xl space-y-6">
            <div class="flex justify-between">
              <span class="text-black text-2xl font-semibold">{{ stats.latest || '-' }}</span>
              <div class="bg-[#071952] rounded-lg p-2"></div>
            </div>
            <span class="text-[#8F8F8F] text-lg">Latest Job</span>
          </div>
        </div>
      </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">            
            <div v-for="job in jobs"
            :key="job.ID" class="bg-white p-4 rounded-xl space-y-6">
             <div class="flex flex-col space-y-3">
                <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
                  <span class="text-xl font-medium">{{ job.JobTitle }}</span>

                  <router-link
                    :to="`/applicants/${job.ID}`"
                    class="px-3 sm:px-4 py-2 bg-white text-purple-600 text-sm sm:text-base font-medium rounded-xl hover:bg-gray-100"
                  >
                    View applicants
                  </router-link>
                </div>

                <div class="flex flex-row items-center gap-4">
                  <div class="px-4 py-2 bg-[#F2F2F2] rounded-md text-[#8F8F8F]">
                    {{ job.Active ? 'Active' : 'Inactive' }}
                  </div>
                  <div class="px-4 py-2 bg-[#F2F2F2] rounded-md text-[#8F8F8F]">
                    Created on: {{ new Date(job.CreatedAt).toLocaleDateString() }}
                  </div>
                </div>

                <!-- 📝 Job Description Section -->
                <div class="bg-[#FAFAFA] border border-gray-200 rounded-lg p-4 text-sm text-gray-700 leading-relaxed">
                  <p class="line-clamp-4">
                    {{ job.Description.substring(0, 150) }}
                  </p>
                </div>
              </div>

              
              <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
                <button @click="copyShareLink(job.ID)"class="text-s font-medium">Copy Share link</button>
                <button
                @click="deleteJob(job.ID)"
                class="px-3 py-1 text-sm bg-red-500 text-white rounded hover:bg-red-600"
                >Delete</button>
              </div>
            </div>
        </div>
    </div>
  </div>

  <!-- Create Job Modal -->
  <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center p-4">
    <!-- Backdrop with blur -->
    <div @click="showModal = false" class="absolute inset-0 bg-black/20 backdrop-blur-sm transition-all"></div>
    
    <!-- Modal Content -->
    <div class="relative bg-white rounded-2xl shadow-xl w-full max-w-lg overflow-hidden transform transition-all mx-2 sm:mx-0">
      <!-- Header -->
      <div class="flex items-center justify-between p-6 pb-2">
        <h2 class="text-2xl font-semibold text-gray-900">Create Job</h2>
        <button @click="showModal = false" class="text-gray-400 hover:text-gray-600 transition p-1">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Form -->
      <!-- BACKEND: Handle form submission to /add-job endpoint -->
      <form @submit.prevent="addJob" class="p-6 pt-2 space-y-5">
        
        <!-- Job Title -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-900">Job Title</label>
          <input 
            v-model="form.job_title" 
            type="text" 
            placeholder="Enter Job Name"
            class="w-full bg-[#f4f4f5] border-transparent focus:bg-white focus:border-blue-500 focus:ring-0 rounded-lg px-4 py-3 text-gray-900 placeholder-gray-400 transition"
          />
        </div>

        <!-- Salary -->
        <div class="space-y-2">
          <label class="block text-sm font-medium text-gray-900">Description</label>
          <div class="flex gap-3">
             <textarea 
              v-model="form.description"
              type="text" 
              placeholder="Enter Salary Range"
              class="flex-1 bg-[#f4f4f5] border-transparent focus:bg-white focus:border-blue-500 focus:ring-0 rounded-lg px-4 py-3 text-gray-900 placeholder-gray-400 transition"
            ></textarea>
          </div>
        </div>
        
        <!-- Submit Button -->
        <div class="pt-4">
            <button type="submit" class="w-full bg-[#0A1B3F] text-white font-medium py-3 rounded-xl hover:bg-[#162a5c] transition shadow-lg shadow-blue-900/10">
                Publish Job
            </button>
        </div>

      </form>
    </div>
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
const showProfileMenu = ref(false);

function logout() {
  localStorage.removeItem("jwt");
  localStorage.removeItem("refresh_token");

  router.push("/login"); 
}

async function deleteJob(jobId) {
  // Authentication check disabled
  const token = localStorage.getItem("jwt");
  if (!token) {
    router.push("/login");
    return;
  }

  if (!confirm("Are you sure you want to delete this job?")) return;

  try {
    const res = await apiRequest(`http://127.0.0.1:8080/delete-job/${jobId}`, {
      method: "DELETE",
    });

    if (!res.ok) {
      alert("Failed to delete job");
      return;
    }
   
    console.log("Mock delete success for:", jobId);

    // Update UI
    jobs.value = jobs.value.filter((j) => j.ID !== jobId);
    alert("Job deleted successfully!");
  } catch (err) {

    alert("Failed to delete job");
    return;
  }
}

async function fetchJobs() {
    // Authentication check disabled
  const token = localStorage.getItem("jwt");
  if (!token) {
    router.push("/login");
    return;
  }

  try {
    const response = await apiRequest("http://127.0.0.1:8080/jobs");

    if (!response.ok) {
      router.push("/login");
      return;
    }

    const data = await response.json();
    jobs.value = data;
    console.log(data)

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
    router.push("/login");
  }
  
}

async function addJob() {
    // Authentication check disabled
  const token = localStorage.getItem("jwt");
  if (!token) {
    router.push("/login");
    return;
  }

  try {
     // Attempt backend call
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
    stats.value.total = data.length;
    stats.value.active = data.filter((j) => j.Active).length;
    stats.value.inactive = data.length - stats.value.active;
    if (data.length > 0) {
      const latestJob = data.reduce((a, b) =>
      new Date(a.CreatedAt) > new Date(b.CreatedAt) ? a : b
      );
      stats.value.latest = latestJob.JobTitle;
    }
    
    
  } catch (error) {
    console.error("Error posting job:", error);
  }
}

const copyShareLink = (jobId) => {
  const link = `http://localhost:5173/apply/${jobId}`; 
  navigator.clipboard.writeText(link).then(() => {
    alert("✅ Share link copied to clipboard!");
  });
}

onMounted(fetchJobs);
</script>
