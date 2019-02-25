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
axios.interceptors.response.use(
  response => {
    const code = response.data.code
    switch (code) {
      case 1001:
        router.push({path: '/login'})
        break
    }
    return response
  }
)

Vue.prototype.$axios = axios

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
