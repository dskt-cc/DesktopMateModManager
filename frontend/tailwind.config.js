/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
            colors: {
                miku: {
                    gray: '#373b3e',
                    light: '#bec8d1',
                    teal: '#86cecb',
                    deep: '#137a7f',
                    pink: '#e12885',
                    aqua: '#80d4e5',
                    waterleaf: '#a1e2e8',
                    padua: '#a6e3d2',
                    ice1: '#b1f1e2',
                    ice2: '#c1f6da',
                },
            },
        },
    },
    plugins: [],
}
