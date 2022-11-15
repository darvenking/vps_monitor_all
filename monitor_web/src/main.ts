import Vue from 'vue'
import App from './App.vue'
import {createPinia, PiniaVuePlugin} from 'pinia'
import router from './router'
import 'font-awesome/css/font-awesome.min.css';
import Notification from 'element-ui/lib/notification';
import 'element-ui/lib/theme-chalk/index.css';
Vue.prototype.$notify = Notification
Vue.use(PiniaVuePlugin)
const pinia = createPinia()
Vue.prototype.$ELEMENT = {size: 'mini', zIndex: 3000}
new Vue({
  render: (h) => h(App),
  pinia,
  router,
}).$mount('#app')
