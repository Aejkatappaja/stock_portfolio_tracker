const BASE_URL = "http://localhost:80/api/txn";

export const API = {
  GET_TRANSACTION: BASE_URL,
  ADD_TRANSACTION: `${BASE_URL}/add`,
  EDIT_TRANSACTION: `${BASE_URL}/edit`,
  DELETE_TRANSACTION: `${BASE_URL}/delete`,
};
