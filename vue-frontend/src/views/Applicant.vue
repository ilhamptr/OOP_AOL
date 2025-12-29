<template>
  <div class="bg-[#F5F7FB] p-6 min-h-screen">
    <div class="max-w-6xl mx-auto space-y-6">

      <!-- Header -->
      <div>
        <h1 class="text-2xl font-bold text-[#0B1F5E]">Job Applications</h1>
        <p class="text-gray-500">Manage and review job applications from candidates</p>
      </div>

      <!-- Applicants List -->
      <div class="bg-white rounded-2xl shadow-sm border border-[#E6ECF7] p-5">

        <div class="flex items-center justify-between mb-5">
          <h2 class="text-lg font-semibold text-gray-800">Applicant List</h2>

          <!-- Top Applicants Control -->
          <div class="flex items-center gap-2">
            <input
              v-model="topNum"
              type="number"
              min="1"
              placeholder="Top Candidate"
              class="w-20 border border-[#E6ECF7] rounded-lg px-3 py-1.5 text-sm
                     focus:ring-2 focus:ring-[#0B1F5E]/30 focus:outline-none"
            />
            <button
              @click="handleTopApplicants"
              class="px-4 py-1.5 bg-[#0B1F5E] text-white text-xs rounded-lg hover:bg-[#09184A] transition"
            >
              Show Top
            </button>
          </div>
        </div>

        <!-- Table -->
        <div class="overflow-x-auto">
          <table class="w-full text-sm text-left border-collapse">
            <thead>
              <tr class="bg-[#F0F3FA] text-gray-700">
                <th class="px-4 py-2">Applicant Name</th>
                <th class="px-4 py-2">Contact</th>
                <th class="px-4 py-2">Resume</th>
                <th class="px-4 py-2">Applied Date</th>
                <th class="px-4 py-2">Actions</th>
              </tr>
            </thead>

            <tbody class="divide-y divide-[#E6ECF7]">
              <tr v-for="(app, index) in applicants" :key="app.ID || index">
                <td class="px-4 py-3 font-medium text-gray-800">
                  {{ app.ApplicantName }}
                  <div class="text-xs text-gray-400">ID: {{ app.ID.slice(0, 8) }}...</div>
                </td>

                <td class="px-4 py-3">
                  <a :href="`mailto:${app.Email}`" class="text-[#0B1F5E] hover:underline">{{ app.Email }}</a>
                  <div class="text-gray-500">{{ app.PhoneNumber }}</div>
                </td>

                <td class="px-4 py-3">
                  <button
                    @click="downloadResume(app.ResumeFile)"
                    class="px-3 py-1.5 bg-[#F0F3FA] rounded-lg text-[#0B1F5E] text-xs font-medium hover:bg-[#E6ECF7]"
                  >
                    Download Resume
                  </button>
                </td>

                <td class="px-4 py-3 text-gray-500">
                  {{ new Date(app.CreatedAt).toLocaleString() }}
                </td>

                <td class="px-4 py-3 flex gap-2">
                  <button
                    @click="reviewApplicant(app.ResumeFile)"
                    class="px-4 py-1.5 bg-[#0B1F5E] text-white text-xs rounded-lg hover:bg-[#09184A]"
                  >
                    Review
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Review Modal -->
      <div
        v-if="showModal"
        class="fixed inset-0 flex items-center justify-center bg-black/40 z-50"
      >
        <div class="bg-white rounded-2xl shadow-lg w-full max-w-2xl p-6 relative border border-[#E6ECF7]">

          <button
            @click="closeModal"
            class="absolute top-3 right-3 text-gray-400 hover:text-gray-600"
          >
            &times;
          </button>

          <h2 class="text-xl font-semibold text-[#0B1F5E] mb-4">Applicant Evaluation</h2>

          <div class="bg-[#F5F7FB] border border-[#E6ECF7] rounded-lg p-4 text-gray-700 leading-relaxed">
            {{ evaluation }}
          </div>

          <div class="mt-6 flex justify-end">
            <button
              @click="closeModal"
              class="px-5 py-2 bg-[#0B1F5E] text-white rounded-lg hover:bg-[#09184A]"
            >
              Close
            </button>
          </div>

        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { apiRequest } from "../lib/api";

const applicants = ref([]);
const evaluation = ref("");
const showModal = ref(false);
const topNum = ref("");

const route = useRoute();
const router = useRouter();

const jobId = route.params.id; 

async function fetchApplicants() {
  const token = localStorage.getItem("jwt");
  if (!token) {
    router.push("/login");
    return;
  }
  

  try {
    const res = await apiRequest(`http://127.0.0.1:8080/view-applicants/${encodeURIComponent(jobId)}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!res.ok) {
        throw new Error("Backend not available or unauthorized");
    }

    const data = await res.json();
    applicants.value = data.map(app => ({
      ID: app.ID || app.id,
      ApplicantName: app.ApplicantName || app.applicant_name,
      Email: app.Email,
      PhoneNumber: app.PhoneNumber,
      ResumeFile: app.ResumeFile || app.resume_file,
      CreatedAt: app.CreatedAt || app.created_at,
    }));
  } catch (err) {
    router.push("/login")
  }
}

async function handleTopApplicants() {
  const token = localStorage.getItem("jwt");
  if (!token) {
    router.push("/login");
    return;
  }

  try {
    const res = await apiRequest(
      `http://127.0.0.1:8080/view-top-applicants/${encodeURIComponent(jobId)}?topNum=${topNum.value}`,
      {
        method: "POST",
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    if (!res.ok) {
        throw new Error("Backend not available or unauthorized");
    }

    const result = await res.json();

    console.log("TOP API RAW:", result);   // keep this for now

    applicants.value = (result?.result?.data || []).map(app => ({
      ID: app.ID || app.id,
      ApplicantName: app.ApplicantName || app.applicant_name,
      Email: app.Email || app.email,
      PhoneNumber: app.PhoneNumber || app.phone,
      ResumeFile: app.ResumeFile || app.resume_file,
      CreatedAt: app.CreatedAt || app.created_at,
    }));
  } catch (err) {
    alert("something went wrong");
    return
  }
}

async function downloadResume(fileId) {
  const token = localStorage.getItem("jwt");
  if (!token) {
    alert("You must be logged in!");
    return;
  }

  try {
    const res = await apiRequest(`http://127.0.0.1:8080/download-resume/${encodeURIComponent(fileId)}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    const data = await res.json();
    const fileUrl = data.url;

    const link = document.createElement("a");
    link.href = fileUrl;
    link.download = "";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

  } catch (err) { 
    console.error("Error downloading resume:", err);
    alert("Download failed.");
   }
}

async function reviewApplicant(resumeFile) {
  const token = localStorage.getItem("jwt");
  if (!token) {
    alert("You must be logged in!");
    return;
  }

  try {
    const res = await fetch(
      `http://127.0.0.1:8080/view-applicant-evaluation/${encodeURIComponent(jobId)}/${encodeURIComponent(resumeFile)}`,
      {
        method: "POST",
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    if (!res.ok) {
      alert("Failed to fetch evaluation");
      return;
    }

    const evalRes = await res.json();
    evaluation.value = evalRes?.result?.response || "No evaluation available.";
   
    showModal.value = true;
  } catch (err) {
    console.error("Error fetching evaluation:", err);
    alert("Something went wrong.");
  }
}

function closeModal() {
  showModal.value = false;
}

onMounted(fetchApplicants);
</script>
