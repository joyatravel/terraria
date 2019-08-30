import Vue from "vue";
import App from "./App.vue";
import router from "./router";

// Do not use store until necessary.
// import store from "./store";

Vue.config.productionTip = false;

new Vue({
  router,
  /* store, */
  render: h => h(App),
}).$mount("#app");
