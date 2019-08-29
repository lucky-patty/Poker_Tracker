import Vue from 'vue';
import App from './App.vue';
import router from './router';
import './registerServiceWorker';
import vuetify from './plugins/vuetify';
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader

Vue.config.productionTip = false;

Vue.use(vuetify,{
  iconfont: 'mdi',
});

new Vue({
  router,
  vuetify,
  render: h => h(App),
}).$mount('#app');
