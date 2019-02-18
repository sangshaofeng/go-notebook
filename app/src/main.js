import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ButtonComponent from './components/Button.vue'
import InputComponent from './components/Input'
import iView from 'iview';
import 'iview/dist/styles/iview.css';

Vue.config.productionTip = false

Vue.use(iView);

Vue.component('ButtonComponent', ButtonComponent)
Vue.component('InputComponent', InputComponent)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
