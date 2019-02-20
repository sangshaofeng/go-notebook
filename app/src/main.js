import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ButtonComponent from './components/Button.vue'
import InputComponent from './components/Input'
import iView from 'iview';
import 'iview/dist/styles/iview.css';
import axios from 'axios'

Vue.config.productionTip = false

Vue.use(iView);

Vue.component('ButtonComponent', ButtonComponent)
Vue.component('InputComponent', InputComponent)

// 响应拦截器
// axios.interceptors.response.use(
//   response => {
//     const status = response.data.status
//     switch (status) {
//       case 500:
//         router.push({path: '/login'})
//     }
//   }
// )

Vue.prototype.$axios = axios

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
