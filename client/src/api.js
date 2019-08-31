// eslint-disable-next-line no-unused-vars
import { create, AxiosInstance } from "axios";

export const baseURL = "/api";

/** @type {AxiosInstance} */
export const client = create({ baseURL });
