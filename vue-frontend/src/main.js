import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './style.css'
// import './assets/main.css'   // ✅ this line is required

const app = createApp(App)
app.use(router)
app.mount('#app')
