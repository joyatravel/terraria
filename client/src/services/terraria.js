import { client, baseURL } from "@/api";

const terraria = {
  async getStatus() {
    const { data } = await client.get("/status");
    return data;
  },

  async downloadWorld() {
    window.open(`${baseURL}/world`);
  },
};

export default terraria;
