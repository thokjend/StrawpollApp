import { createRouter, createWebHistory } from "vue-router";

import LoginPage from "../pages/Login.vue";
import CreatePage from "../pages/Create.vue";
import ViewPollPage from "../pages/ViewPoll.vue";
import ViewResultPage from "../pages/ViewResults.vue";
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
  {
    path: "/create",
    name: "Create",
    component: CreatePage,
  },
  {
    path: "/poll/:id",
    name: "ViewPoll",
    component: ViewPollPage,
    props: true,
  },
  {
    path: "/poll/:id/result",
    name: "ViewResult",
    component: ViewResultPage,
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
