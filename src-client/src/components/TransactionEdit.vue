<template>
  <div class="flex min-h-screen items-center justify-center">
    <div class="w-full max-w-md rounded-lg bg-white p-8 shadow-lg">
      <div class="mb-2">
        <h1 class="text-2xl font-semibold text-gray-700">Edit Transaction</h1>
      </div>

      <section class="mb-2 flex h-10 w-full items-start rounded-lg">
        <ul class="flex-1 rounded-lg bg-red-600" v-if="err != ''">
          <li class="px-4 py-2 font-semibold text-black">{{ err }}</li>
        </ul>
      </section>

      <div class="space-y-4">
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
          @click="editTransaction"
          class="w-full rounded-lg bg-blue-500 py-2 font-semibold text-white transition duration-300 hover:bg-blue-600"
        >
          Save
        </button>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { TransactionModel } from "../definitions";
import { API } from "../utils/api";
import axios from "axios";
import { defineComponent } from "vue";

export default defineComponent({
  name: "transactionEdit",
  data() {
    return {
      err: null,
      model: {
        transaction: {
          id: 0,
          type: "",
          ticker: "",
          volume: "",
          price: "",
          date: "",
        },
      },
    };
  },
  methods: {
    async getTransactionData(id: number) {
      try {
        const res = await axios.get(`${API.GET_TRANSACTION}/${id}`);
        this.model.transaction = res.data.transaction;
        console.log(this.model.transaction);
      } catch (e: any) {
        console.error("Error fetching transaction", e);
      }
    },

    async editTransaction() {
      try {
        const res = await axios.post(API.EDIT_TRANSACTION, {
          id: Number(this.model.transaction.id),
          type: this.model.transaction.type,
          ticker: this.model.transaction.ticker,
          volume: Number(this.model.transaction.volume),
          price: Number(this.model.transaction.price),
          date: this.model.transaction.date,
        });

        console.log(res);

        this.model.transaction = {
          id: 0,
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
  mounted() {
    this.getTransactionData(Number(this.$route.params.id));
  },
});
</script>
