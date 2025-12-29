import { createRouter, createWebHistory } from 'vue-router'

import About from '../views/About.vue'
// import Login from '../views/Login.vue'
import Main from '../views/Main.vue'
// import Apply from '../views/Apply.vue'


// import Main from '../views/Main.vue'   // create this later

const routes = [
    { path: '/', name: 'Home', component: Main }, // Using Main component for root
    { path: '/about', name: 'About', component: About },
    { path: "/login", name: "Login", component: () => import("../views/Login.vue")},
    { path: "/main", redirect: '/' }, // Redirect /main to /
    { path: "/apply/:id", name: "Apply", component: () => import("../views/Apply.vue")},
    { path: "/applicants/:id", name: "Applicants", component: () => import("../views/Applicant.vue")},
    {path: "/auth/callback",name: "callback",component: () => import("../views/AuthCallback.vue")}

]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
