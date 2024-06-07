/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['views/*.html',],
  theme: {
    extend: {
      colors: {
        'lime': '#8FBC8F',
        '7580ff': '#7580ff',
        'squeeze': '#5423e7',
      },
      fontFamily: {
        poppins: ['Poppins'],
        montserrat: ['Montserrat'],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "luxury", "sunset"],
  },
}
