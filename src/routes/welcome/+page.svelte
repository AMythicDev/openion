<script lang="ts">
	import * as Tabs from '$lib/components/ui/tabs/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import { Button } from '$lib/components/ui/button';
	import { Checkbox } from '$lib/components/ui/checkbox/';
	import type { PageProps } from '../$types.js';
	import axios from 'axios';
	import { API_URL } from '$lib';
	import { user } from '$lib/user.svelte.js';
	import { goto } from '$app/navigation';
	import * as Alert from '$lib/components/ui/alert/index.js';

	const { data }: PageProps = $props();
	const id = $props.id();

	let lEmail = $state();
	let lPassword = $state();
	let lShowpw = $state(false);

	let sName = $state();
	let sEmail = $state();
	let sPassword = $state();
	let sCPassword = $state();
	let sShowpw = $state(false);

	let err = $state<string | null>(null);

	async function login() {
		try {
			const resp = await axios.post(
				`${API_URL}/login`,
				{
					email: lEmail,
					password: lPassword
				},
				{
					withCredentials: true
				}
			);
			user.current = resp.data;
			goto('/');
		} catch (error) {
			if (error.status == 401) {
				err = error.response.data.error;
			} else {
				console.log(error);
			}
		}
	}

	async function signup() {
		try {
			const resp = await axios.post(
				`${API_URL}/newuser`,
				{
					name: sName,
					email: sEmail,
					password: sPassword
				},
				{
					withCredentials: true
				}
			);
			user.current = resp.data;
			goto('/');
		} catch (error) {
			console.log(error);
		}
	}
</script>

<div class="flex min-h-[36] flex-col items-center">
	{#if err}
		<Alert.Root variant="destructive" class="w-[36rem] bg-destructive/10">
			<Alert.Title>Unable to process your payment.</Alert.Title>
		</Alert.Root>
	{/if}
	<Tabs.Root value={data.form} class="min-w-[36rem] bg-white">
		<Tabs.List class="w-full">
			<Tabs.Trigger value="login">Login</Tabs.Trigger>
			<Tabs.Trigger value="signup">Sign Up</Tabs.Trigger>
		</Tabs.List>
		<Tabs.Content value="login" class="px-8 py-2">
			<h1 class="mb-4 text-2xl">Welcome Back</h1>
			<form onsubmit={login}>
				<Label for="lemail-{id}" class="mt-4 mb-1">Enter your email</Label>
				<Input
					type="email"
					id="lemail-{id}"
					bind:value={lEmail}
					class="focus-visible:border-primary focus-visible:ring-primary"
				/>
				<Label for="lpw-{id}" class="mt-4 mb-1">Enter your password</Label>
				<Input
					type={lShowpw ? 'text' : 'password'}
					id="lpw-{id}"
					bind:value={lPassword}
					class="focus-visible:border-primary focus-visible:ring-primary"
				/>
				<div class="mt-2 flex items-center gap-1">
					<Checkbox id="lshowpw-{id}" class="inline border-black" bind:checked={lShowpw} />
					<Label for="lshowpw-{id}" class="inline">Show password</Label>
				</div>
				<div class="mt-4 flex w-full justify-end">
					<Button type="submit">Login</Button>
				</div>
			</form>
		</Tabs.Content>
		<Tabs.Content value="signup" class="px-8 py-2">
			<h1 class="mb-4 text-2xl">Get Started on Openion</h1>
			<form onsubmit={signup}>
				<Label for="sname-{id}" class="mt-4 mb-1">Enter your full name</Label>
				<Input
					type="text"
					id="sname-{id}"
					bind:value={sName}
					class="focus-visible:border-secondary focus-visible:ring-secondary"
				/>
				<Label for="semail-{id}" class="mt-4 mb-1">Enter your email</Label>
				<Input
					type="email"
					id="semail-{id}"
					bind:value={sEmail}
					class="focus-visible:boring-secondary focus-visible:ring-secondary"
				/>
				<Label for="spw-{id}" class="mt-4 mb-1">Enter a password</Label>
				<Input
					type={lShowpw ? 'text' : 'password'}
					id="spw-{id}"
					bind:value={sPassword}
					class="focus-visible:boring-secondary focus-visible:ring-secondary"
				/>
				<Label for="scpw-{id}" class="mt-4 mb-1">Confirm password</Label>
				<Input
					type={lShowpw ? 'text' : 'password'}
					id="scpw-{id}"
					bind:value={sCPassword}
					class="focus-visible:boring-secondary focus-visible:ring-secondary"
				/>
				<div class="mt-2 flex items-center gap-1">
					<Checkbox
						id="sshowpw-{id}"
						class="inline border-black data-[state=checked]:bg-secondary dark:data-[state=checked]:bg-secondary"
						bind:checked={sShowpw}
					/>
					<Label for="sshowpw-{id}" class="inline">Show password</Label>
				</div>
				<div class="mt-4 flex w-full justify-end">
					<Button type="submit" class="bg-secondary hover:bg-secondary/90">Sign Up</Button>
				</div>
			</form>
		</Tabs.Content>
	</Tabs.Root>
</div>
