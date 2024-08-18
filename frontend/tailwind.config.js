/** @type {import('tailwindcss').Config} */
export default {
  darkMode: "class",
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
    "./public/fonts/**/*.{ttf,woff,woff2}",
  ],
  theme: {
    extend: {
      fontFamily: {
        JetBrainsMonoBold: ['JetBrainsMonoBold', 'JetBrainsMonoExtraBold', 'sans-serif'],
        JetBrainsMonoExtraBold: ['JetBrainsMonoExtraBold', 'sans-serif'],
      },
    },
  },
  plugins: [],
}