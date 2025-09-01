<template>
  <div class="bg-violet-50 p-6 min-h-screen">
    <div class="max-w-6xl mx-auto space-y-6">

      <!-- Header -->
      <div>
        <h1 class="text-2xl font-bold text-violet-700">Job Applications</h1>
        <p class="text-gray-600">Manage and review job applications from candidates</p>
      </div>

      <!-- Applicants List -->
      <div class="bg-white rounded-lg shadow p-4">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-semibold text-gray-800">Applicant List</h2>

          <!-- Top Applicants Control -->
          <div class="flex items-center gap-2">
            <input
              v-model="topNum"
              type="number"
              min="1"
              placeholder="Top N"
              class="w-20 border border-gray-300 rounded px-2 py-1 text-sm
                     focus:ring-violet-500 focus:border-violet-500"
            />
            <button
              @click="handleTopApplicants"
              class="px-3 py-1 bg-violet-600 text-white text-xs rounded hover:bg-violet-700"
            >
              Show Top
            </button>
          </div>
        </div>

        <!-- Table -->
        <div class="overflow-x-auto">
          <table class="w-full text-sm text-left border-collapse">
            <thead>
              <tr class="bg-violet-100 text-gray-700">
                <th class="px-4 py-2">Applicant Name</th>
                <th class="px-4 py-2">Contact</th>
                <th class="px-4 py-2">Resume</th>
                <th class="px-4 py-2">Applied Date</th>
                <th class="px-4 py-2">Actions</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="app in applicants" :key="app.ID">
                <td class="px-4 py-3 font-medium text-gray-800">
                  {{ app.ApplicantName }}
                  <br />
                  <span class="text-xs text-gray-500">ID: {{ app.ID.slice(0, 8) }}...</span>
                </td>
                <td class="px-4 py-3">
                  <a :href="`mailto:${app.Email}`" class="text-violet-600 hover:underline">{{ app.Email }}</a>
                  <br />
                  <a :href="`tel:${app.PhoneNumber}`" class="text-gray-600">{{ app.PhoneNumber }}</a>
                </td>
                <td class="px-4 py-3">
                  <button
                    @click="downloadResume(app.ResumeFile)"
                    class="px-2 py-1 bg-gray-100 rounded text-violet-700 text-xs font-medium hover:bg-violet-200"
                  >
                    Download Resume
                  </button>
                </td>
                <td class="px-4 py-3 text-gray-600">
                  {{ new Date(app.CreatedAt).toLocaleString() }}
                </td>
                <td class="px-4 py-3 flex gap-2">
                  <button
                    @click="reviewApplicant(app.ResumeFile)"
                    class="px-3 py-1 bg-violet-600 text-white text-xs rounded hover:bg-violet-700"
                  >
                    Review
                  </button>
                  <button
                    class="px-3 py-1 bg-gray-200 text-gray-700 text-xs rounded hover:bg-gray-300"
                  >
                    Contact
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
        class="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50"
      >
        <div class="bg-white rounded-lg shadow-lg w-full max-w-2xl p-6 relative">
          <button
            @click="closeModal"
            class="absolute top-3 right-3 text-gray-500 hover:text-gray-700"
          >
            &times;
          </button>

          <h2 class="text-xl font-semibold text-violet-700 mb-4">Applicant Evaluation</h2>
          <div class="bg-violet-50 border border-violet-200 rounded p-4 text-gray-700 leading-relaxed">
            {{ evaluation }}
          </div>

          <div class="mt-6 flex justify-end">
            <button
              @click="closeModal"
              class="px-4 py-2 bg-violet-600 text-white rounded hover:bg-violet-700"
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

const applicants = ref([]);
const evaluation = ref("");
const showModal = ref(false);
const topNum = ref("");

const route = useRoute();
const router = useRouter();

const jobId = route.params.id; // <-- now using route param (not query)

async function fetchApplicants() {
  const token = localStorage.getItem("jwt");
  if (!token) {
    router.push("/login");
    return;
  }

  try {
    const res = await fetch(`http://127.0.0.1:8080/view-applicants/${encodeURIComponent(jobId)}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!res.ok) {
      router.push("/login");
      return;
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
    console.error("Error fetching applicants:", err);
    alert("Something went wrong.");
  }
}

async function handleTopApplicants() {
  const token = localStorage.getItem("jwt");
  if (!token) {
    router.push("/login");
    return;
  }

  try {
    const res = await fetch(
      `http://127.0.0.1:8080/view-top-applicants/${encodeURIComponent(jobId)}?topNum=${topNum.value}`,
      {
        method: "POST", // change to GET if backend expects GET
        headers: { Authorization: `Bearer ${token}` },
      }
    );

    if (!res.ok) {
      alert("Failed to fetch top applicants");
      return;
    }

    const result = await res.json();
    applicants.value = (result?.result?.data || []).map(app => ({
      ID: app.ID || app.id,
      ApplicantName: app.ApplicantName || app.applicant_name,
      Email: app.Email,
      PhoneNumber: app.PhoneNumber,
      ResumeFile: app.ResumeFile || app.resume_file,
      CreatedAt: app.CreatedAt || app.created_at,
    }));
  } catch (err) {
    console.error("Error fetching top applicants:", err);
    alert("Something went wrong.");
  }
}

async function downloadResume(fileId) {
  const token = localStorage.getItem("jwt");
  if (!token) {
    alert("You must be logged in!");
    return;
  }

  try {
    const res = await fetch(`http://127.0.0.1:8080/download-resume/${encodeURIComponent(fileId)}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!res.ok) {
      alert("Failed to fetch resume");
      return;
    }

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
