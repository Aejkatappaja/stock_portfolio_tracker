import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "../components/Dashboard.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "dashboard",
      component: Dashboard,
    },
    // {
    //   path: "/transactions/create",
    //   name: "transactionCreate",
    //   component: TransactionCreate,
    // },
    // {
    //   path: "/transactions/:id/edit",
    //   name: "transactionEdit",
    //   component: TransactionEdit,
    // },
  ],
});
export default router;
