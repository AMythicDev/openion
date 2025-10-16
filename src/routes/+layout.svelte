<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import { user } from '$lib/user.svelte';
	import axios from 'axios';
	import { API_URL } from '$lib';

	let { children } = $props();

	async function fetchUser() {
		if (!user.current) {
			const resp = await axios.get(`${API_URL}/getuser`, { withCredentials: true });
      if (resp.data)
        user.current = resp.data;
      else 
        throw "unauthenticated"
		}
	}
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<nav class="mb-6 flex h-14 items-center justify-between bg-white px-12 shadow-sm">
	<a
		class="bg-linear-to-r from-primary to-secondary bg-clip-text font-brand text-3xl text-transparent"
		href="/"
	>
		Openion
	</a>
	{#await fetchUser() catch}
		<div class="flex gap-3">
			<div
				class="flex h-10 items-center justify-center rounded-xl bg-gradient-to-r from-primary to-secondary p-0.5"
			>
				<a
					href="/welcome?login"
					class="flex h-full w-full items-center justify-center rounded-xl bg-white px-3 transition-colors outline-none hover:bg-gradient-to-r hover:from-primary hover:to-secondary hover:text-white"
				>
					Login
				</a>
			</div>
			<div
				class="flex h-10 items-center justify-center rounded-xl bg-gradient-to-r from-primary to-secondary p-0.5"
			>
				<a
					href="/welcome?signup"
					class="flex h-full w-full items-center justify-center rounded-xl bg-gradient-to-r from-primary to-secondary px-3 text-white transition-colors outline-none hover:bg-white hover:bg-none hover:text-black"
				>
					Sign Up
				</a>
			</div>
		</div>
	{/await}
</nav>
{@render children?.()}
