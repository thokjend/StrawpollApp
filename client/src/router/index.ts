import { createRouter, createWebHistory } from "vue-router";

import LoginPage from "../pages/Login.vue";
import CreatePage from "../pages/Create.vue";

const routes = [
  {
    path: "/",
    name: "Login",
    component: LoginPage,
  },
  {
    path: "/create",
    name: "Create",
    component: CreatePage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
