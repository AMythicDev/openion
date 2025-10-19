<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import { user } from '$lib/user.svelte';
	import axios from 'axios';
	import { API_URL, shortenName } from '$lib';
	import * as Avatar from '$lib/components/ui/avatar/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	let { children } = $props();

	async function fetchUser() {
		if (!user.current) {
			const resp = await axios.get(`${API_URL}/getuser`, { withCredentials: true });
			if (resp.data) {
				user.current = resp.data;
				return resp.data;
			} else throw 'unauthenticated';
		}
	}

  async function logout() {
			await axios.get(`${API_URL}/logout`, { withCredentials: true });
      user.current = null;
  }
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<nav class="mb-6 flex h-14 items-center justify-between bg-white px-4 shadow-sm md:px-12">
	<a
		class="bg-linear-to-r from-primary to-secondary bg-clip-text font-brand text-3xl text-transparent"
		href="/"
	>
		Openion
	</a>
	{#await fetchUser() then}
		<DropdownMenu.Root>
			<DropdownMenu.Trigger class="flex gap-2 items-center">
				<Avatar.Root class="h-10 w-10 shadow-sm">
					<Avatar.Image src={user.current.Avatar} alt="user image" />
					<Avatar.Fallback>{shortenName(user.current.Name)}</Avatar.Fallback>
				</Avatar.Root>
        {user.current.Name}
			</DropdownMenu.Trigger>
			<DropdownMenu.Content>
				<DropdownMenu.Group>
					<DropdownMenu.Label>Account</DropdownMenu.Label>
					<DropdownMenu.Item variant="destructive" onclick={logout}>Logout</DropdownMenu.Item>
				</DropdownMenu.Group>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	{:catch}
		<div class="flex gap-3">
			<div
				class="hidden items-center justify-center rounded-xl bg-gradient-to-r from-primary to-secondary p-0.5 md:flex md:h-10"
			>
				<a
					href="/welcome?login"
					class="flex h-full w-full items-center justify-center rounded-xl bg-white px-3 transition-colors outline-none hover:bg-gradient-to-r hover:from-primary hover:to-secondary hover:text-white"
				>
					Login
				</a>
			</div>
			<div
				class="flex h-8 w-22 items-center justify-center rounded-xl bg-gradient-to-r from-primary to-secondary p-0.5 md:h-10"
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

