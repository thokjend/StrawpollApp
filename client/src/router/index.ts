import { createRouter, createWebHistory } from "vue-router";

import LoginPage from "../pages/Login.vue";
import CreatePage from "../pages/Create.vue";
import ViewPollPage from "../pages/ViewPoll.vue";

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
  {
    path: "/poll/:id",
    name: "View",
    component: ViewPollPage,
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
