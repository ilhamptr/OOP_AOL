import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'
// import Login from '../views/Login.vue'
import Main from '../views/Main.vue'
// import Apply from '../views/Apply.vue'


// import Main from '../views/Main.vue'   // create this later

const routes = [
    { path: '/', name: 'Home', component: Home },
    { path: '/about', name: 'About', component: About },
    { path: "/login", name: "Login", component: () => import("../views/Login.vue")},
    { path: "/main", name: "Main", component: Main },
    {path: "/apply/:id",name: "Apply",component: () => import("../views/Apply.vue")},
    {path: "/applicants/:id",name: "Applicants",component: () => import("../views/Applicant.vue")},
    {path: "/forgot-password",name: "Forgot-Password",component: () => import("../views/ForgotPassword.vue")},
    {path: "/signup",name: "Signup",component: () => import("../views/signup.vue")}
    //   { path: '/main', name: 'Main', component: Main },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
