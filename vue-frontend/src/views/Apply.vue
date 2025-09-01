<template>
  <div class="bg-purple-50 min-h-screen">
    <!-- Header -->
    <div class="text-center py-8">
      <h1 class="text-3xl font-bold text-purple-700">Join Our Team</h1>
      <p class="text-gray-600">We're looking for talented individuals to help us grow</p>
    </div>

    <div class="max-w-6xl mx-auto grid grid-cols-1 md:grid-cols-3 gap-6 px-4">
      <!-- Job Details -->
      <div class="col-span-2 space-y-6">
        <div class="bg-white shadow rounded-lg overflow-hidden">
          <div v-html="jobHeader" class="bg-purple-600 text-white p-6"></div>
          <div v-html="jobDescription" class="p-6 space-y-3 text-gray-700"></div>
        </div>
      </div>

      <!-- Application Form -->
      <div class="bg-white shadow rounded-lg p-6">
        <h3 class="text-lg font-semibold text-purple-700 mb-2">Apply for this Position</h3>
        <p class="text-sm text-gray-500 mb-4">
          Fill out the form below to submit your application
        </p>

        <form @submit.prevent="submitApplication" class="space-y-3">
          <input v-model="form.name" type="text" placeholder="Full Name *" class="w-full border rounded px-3 py-2" required />
          <input v-model="form.email" type="email" placeholder="Email Address *" class="w-full border rounded px-3 py-2" required />
          <input v-model="form.phone_number" type="tel" placeholder="Phone Number *" class="w-full border rounded px-3 py-2" required />
          <input @change="handleFileUpload" type="file" accept=".pdf,.doc,.docx" class="w-full text-sm text-gray-600" required />
          <button type="submit" class="w-full bg-purple-600 text-white py-2 rounded hover:bg-purple-700">
            Submit Application
          </button>
        </form>

        <div v-if="formMessage" class="mt-4 text-sm" v-html="formMessage"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "JobDetails",
  data() {
    return {
      jobHeader: "",
      jobDescription: "",
      form: {
        name: "",
        email: "",
        phone_number: "",
        file: null,
      },
      formMessage: "",
    };
  },
  methods: {
    async fetchJob() {
      try {
        const jobId = this.$route.params.id; // ✅ get from URL param
        const res = await fetch(`http://127.0.0.1:8080/job-details/${jobId}`);
        if (!res.ok) throw new Error("Failed to fetch job details");

        const job = await res.json();

        this.jobHeader = `
          <h2 class="text-2xl font-bold">${job.JobTitle}</h2>
          <div class="flex items-center gap-4 text-sm mt-2">
            <span>📅 Posted ${new Date(job.CreatedAt).toDateString()}</span>
          </div>
        `;

        this.jobDescription = `<pre class="whitespace-pre-line">${job.Description}</pre>`;
      } catch (err) {
        console.error(err);
        this.jobDescription = "<p class='text-red-600'>Error loading job details.</p>";
      }
    },
    handleFileUpload(event) {
      this.form.file = event.target.files[0];
    },
    async submitApplication() {
      const jobId = this.$route.params.id; // ✅ still from URL param
      const formData = new FormData();
      formData.append("name", this.form.name);
      formData.append("email", this.form.email);
      formData.append("phone_number", this.form.phone_number);
      formData.append("file", this.form.file);
      formData.append("job_id", jobId);

      try {
        const response = await fetch(`http://127.0.0.1:8080/apply/${jobId}`, {
          method: "POST",
          body: formData,
        });

        if (!response.ok) throw new Error("Failed to submit application");

        this.formMessage = `<p class="text-green-600">✅ Application submitted successfully!</p>`;
        this.form = { name: "", email: "", phone_number: "", file: null };
      } catch (err) {
        console.error(err);
        this.formMessage = `<p class="text-red-600">❌ Failed to submit application.</p>`;
      }
    },
  },
  mounted() {
    this.fetchJob();
  },
};
</script>
