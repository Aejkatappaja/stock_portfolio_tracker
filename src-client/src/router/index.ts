import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "dashboard",
      component: () => import("../components/Dashboard.vue"),
    },
    {
      path: "/transactions/create",
      name: "transactionCreate",
      component: () => import("../components/TransactionCreate.vue"),
    },
    {
      path: "/transaction/:id/edit",
      name: "transactionEdit",
      component: () => import("../components/TransactionEdit.vue"),
    },
  ],
});
export default router;
