import { defineConfig } from 'vitest/config';
import { sveltekit } from '@sveltejs/kit/vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
		  '/api': 'http://localhost:1145'
		}
	},
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
