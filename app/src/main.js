import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ButtonComponent from './components/Button.vue'
import InputCoponent from './components/Input'

Vue.config.productionTip = false

Vue.component('ButtonComponent', ButtonComponent)
Vue.component('InputCoponent', InputCoponent)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
