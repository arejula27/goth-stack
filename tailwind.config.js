/** @type {import('tailwindcss').Config} */

const colors = require('tailwindcss/colors')
module.exports = {
  content: ["./views/**/*.{html,js}"],
  theme: {
    extend: {
      ...colors
    },
  },
  plugins: [],
}