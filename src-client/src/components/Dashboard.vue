<template>
  <main class="min-h-screen bg-gray-100 pb-6">
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
          <tbody v-if="data">
            <tr
              v-for="t in data.transactions"
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
        <button @click="getAllTransactions" class="text-black">test</button>
      </div>
    </div>
  </main>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";
import { ref, onMounted } from "vue";
import { getTransactions, deleteTransaction } from "../utils/get-transactions";
import type { Transactions } from "../utils/get-transactions";
import { API } from "../utils/api";

const rows: string[] = [
  "ID",
  "Type",
  "Ticker",
  "Volume",
  "Price",
  "Date",
  "Actions",
];

const data = ref<Transactions | null>(null);

async function getAllTransactions() {
  await getTransactions<Transactions>(API.GET_TRANSACTION, data);
}

function deleteTransactionById(id: number) {
  if (confirm("Are you sure you want to remove this transaction ?"))
    deleteTransaction(API.DELETE_TRANSACTION, id);
  getAllTransactions();
}

onMounted(getAllTransactions);
</script>
