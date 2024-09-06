<template>
  <div class="flex min-h-screen items-center justify-center">
    <div class="w-full max-w-md rounded-lg bg-white p-8 shadow-lg">
      <div class="mb-2">
        <h1 class="text-2xl font-semibold text-gray-700">Add Transactions</h1>
      </div>

      <section class="mb-2 flex w-full items-start rounded-lg">
        <ul class="flex-1 rounded-lg bg-red-600" v-if="err != null">
          <li class="px-4 py-2 font-semibold text-black">{{ err }}</li>
        </ul>
      </section>

      <form @submit.prevent="saveTransaction" class="space-y-4">
        <div v-for="r in rows" :key="r">
          <label :for="r" class="mb-1 block font-medium text-gray-600">{{
            r
          }}</label>
          <input
            :id="r"
            type="text"
            autocomplete="off"
            class="w-full rounded-lg border border-gray-300 px-4 py-2 text-black focus:outline-none focus:ring-2 focus:ring-blue-400"
            v-model="
              model.transaction[r.toLowerCase() as keyof TransactionModel]
            "
          />
        </div>

        <button
          type="submit"
          class="w-full rounded-lg bg-blue-500 py-2 font-semibold text-white transition duration-300 hover:bg-blue-600"
        >
          Save
        </button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import axios from "axios";
import { API } from "../utils/api";
import { TransactionModel } from "../definitions";

export default defineComponent({
  name: "transactionCreate",
  data() {
    return {
      err: null as string | null,
      model: {
        transaction: {} as TransactionModel,
      },
    };
  },
  methods: {
    async saveTransaction() {
      this.err = null; // Reset the error state before submission
      try {
        const res = await axios.post(API.ADD_TRANSACTION, {
          type: this.model.transaction.type,
          ticker: this.model.transaction.ticker,
          volume: Number(this.model.transaction.volume),
          price: Number(this.model.transaction.price),
          date: this.model.transaction.date,
        });
        console.log(res);

        this.model.transaction = {
          type: "",
          ticker: "",
          volume: "",
          price: "",
          date: "",
        };

        this.$router.push("/");
      } catch (e: any) {
        console.log(e);
        if (e.response?.status === 400 || e.response?.status === 500) {
          console.log(e.response.data.error);
          this.err = e.response.data.error;
        }
      }
    },
  },
  computed: {
    rows() {
      return ["Type", "Ticker", "Volume", "Price", "Date"];
    },
  },
});
</script>
