import { Ref } from "vue";

export interface Transaction {
  id: number;
  type: string;
  ticker: string;
  volume: number;
  price: number;
  date: string;
}

export interface Transactions {
  transactions: Transaction[];
}

export async function useFetch<T>(
  url: string,
  params: RequestInit = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    body: null,
  },
): Promise<T | null> {
  try {
    const response = await fetch(url, params);
    if (!response.ok) throw new Error(`HTTP error! Status: ${response.status}`);
    const data: T = await response.json();
    console.log(data);
    return data;
  } catch (error) {
    console.error("Failed to fetch data", error);
    return null;
  }
}

export async function getTransactions<T>(url: string, ref: Ref<T | null>) {
  try {
    const data = await useFetch<T>(url);
    if (data) ref.value = data;
  } catch (error) {
    console.error(error);
    return null;
  }
}

export async function deleteTransaction(url: string, id: number) {
  try {
    const data = await useFetch(url, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id }),
    });
    if (data) return data;
  } catch (error) {
    console.error(error);
  }
}
