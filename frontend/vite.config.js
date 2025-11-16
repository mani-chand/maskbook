import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';

// https://vite.dev/config/
export default defineConfig({
    plugins: [tailwindcss(), svelte()], server: {
        proxy: {
            // This proxies all requests starting with /api
            '/api': {
                // Your Go backend's URL
                target: 'http://localhost:8080',

                // This is crucial for the Go server to accept the request
                changeOrigin: true,

                // Optional: rewrite the path if your Go API doesn't have /api
                // rewrite: (path) => path.replace(/^\/api/, ''), 
            }
        }
    }
});
