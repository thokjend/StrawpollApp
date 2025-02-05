import { createRouter, createWebHistory } from "vue-router";

import LoginPage from "../pages/Login.vue";
import MainPage from "../pages/Main.vue";

const routes = [
  {
    path: "/",
    name: "Login",
    component: LoginPage,
  },
  {
    path: "/main",
    name: "Main",
    component: MainPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
