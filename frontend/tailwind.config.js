/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "./app/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: "#1DBEB4",
          50: "#A2F1EC",
          100: "#90EEE8",
          200: "#6DE9E1",
          300: "#49E3DA",
          400: "#26DED2",
          500: "#1DBEB4",
          600: "#168D86",
          700: "#0E5D58",
          800: "#072C2A",
          900: "#000000",
        },
        secondary: {
          DEFAULT: "#383B58",
          50: "#9093B8",
          100: "#8387B0",
          200: "#6A6FA0",
          300: "#585C8A",
          400: "#484C71",
          500: "#383B58",
          600: "#222436",
          700: "#0C0D13",
          800: "#000000",
          900: "#000000",
          950: "#000000",
        },
        ivory: {
          DEFAULT: "#F3F1ED",
          50: "#FFFFFF",
          100: "#FFFFFF",
          200: "#FFFFFF",
          300: "#FFFFFF",
          400: "#FFFFFF",
          500: "#f9f9f9",
          600: "#DDD7CB",
          700: "#C6BDAA",
          800: "#B0A288",
          900: "#998866",
          950: "#89795B",
        },
      },
    },
  },
  plugins: [],
};
