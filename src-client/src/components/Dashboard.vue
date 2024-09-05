<template>
  <main class="min-h-screen pb-6">
    <div class="mx-auto max-w-7xl rounded-lg bg-white p-6 shadow-lg">
      <div class="mb-6 flex items-center justify-between">
        <h4 class="text-xl font-semibold text-gray-700">Transactions</h4>
        <RouterLink
          to="/transactions/create"
          class="rounded-lg bg-blue-500 px-4 py-2 font-semibold text-white shadow-md hover:bg-blue-600"
          >Add Transaction</RouterLink
        >
      </div>
      <div class="overflow-x-auto">
        <table class="min-w-full rounded-lg border border-gray-300 bg-white">
          <thead>
            <tr
              class="bg-gray-200 text-sm uppercase leading-normal text-gray-700"
            >
              <th v-for="r in rows" :key="r" class="px-6 py-3 text-left">
                {{ r }}
              </th>
            </tr>
          </thead>
          <tbody v-if="transactions.length >= 1">
            <tr
              v-for="t in transactions"
              :key="t.id"
              class="border-b border-gray-300 hover:bg-gray-50"
            >
              <td
                v-for="(value, index) in t"
                :key="index"
                class="px-6 py-3 text-black"
              >
                {{ value }}
              </td>
              <td class="flex space-x-2 px-6 py-3">
                <RouterLink
                  :to="{ path: `/transaction/${t.id}/edit` }"
                  class="rounded-lg bg-green-500 px-4 py-2 font-semibold text-white shadow-lg hover:bg-green-600"
                  >EDIT</RouterLink
                >
                <button
                  @click="deleteTransactionById(t.id)"
                  type="button"
                  class="rounded-lg bg-red-500 px-4 py-2 font-semibold text-white shadow-lg hover:bg-red-600"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { API } from "../utils/api";
import axios from "axios";

interface Transaction {
  id: number;
  type: string;
  ticker: string;
  volume: number;
  price: number;
  date: string;
}

export default defineComponent({
  name: "dashboard",

  data() {
    return { transactions: [] as Transaction[] };
  },

  methods: {
    async getTransactions() {
      axios
        .get(API.GET_TRANSACTION)
        .then((res) => {
          this.transactions = res.data.transactions;
        })
        .catch((e) => {
          console.error("Error fetching transactions", e);
        });
    },

    async deleteTransactionById(id: number) {
      axios
        .delete(API.DELETE_TRANSACTION, { data: { id: Number(id) } })
        .then((res) => {
          alert(res.data.message);
          this.getTransactions();
        })
        .catch((e) => {
          if (e.response) {
            if (e.response.status == 400) alert(e.response.data.error);
            if (e.response.status == 500) alert(e.response.data.error);
          }
        });
    },
  },

  computed: {
    rows() {
      return ["ID", "Type", "Ticker", "Volume", "Price", "Date", "Actions"];
    },
  },
  mounted() {
    this.getTransactions();
  },
});
</script>
