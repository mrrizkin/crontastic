/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./resources/views/**/*.{html,templ,js,ts}",
    "./resources/js/**/*.{veu,js,jsx,ts,tsx}",
    "node_modules/preline/dist/*.js",
  ],
  theme: {
    extend: {},
  },
  plugins: [require("preline/plugin")],
};
