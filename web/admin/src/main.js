import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'ant-design-vue/dist/reset.css'
import axios from 'axios'
axios.defaults.baseURL='http://localhost:3000/api/v1'
Vue.prototype.$http=axios
Vue.config.productionTip=false

Vue.use(Button)

createApp(App).use(router).mount('#app')
