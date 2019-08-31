const dotenv = require("dotenv");
const path = require("path");

// Load environment variables.
[".env", ".env.local", ".env.development", ".env.development.local"].forEach(
  file => dotenv.config({ path: path.resolve(process.cwd(), file) })
);

const { VUE_APP_API_BASE_URL } = process.env;

module.exports = {
  devServer: {
    proxy: {
      "^/api": {
        target: VUE_APP_API_BASE_URL,
        pathRewrite: { "^/api/": "/" },
      },
    },
  },
};
