import Axios from "axios";
import { setupCache } from "axios-cache-interceptor";

const instance = Axios.create({
  headers: {
    common: {
      "X-Requested-With": "XMLHttpRequest",
    },
  },
});
const axios = setupCache(instance);

export default axios;
