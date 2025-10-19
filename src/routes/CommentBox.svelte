<script lang="ts">
	let { class: className = null, parent = null } = $props();
	import Comment from './Comment.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { API_URL } from '$lib';
	import { user } from '$lib/user.svelte';
	import axios from 'axios';

	let loadedComments: any = $state([]);
	let page = 1;
	let moreComments = $state(true);
	const PAGESIZE: number = 20;
	let replyText = $state('');

	async function getThreadComment() {
		if (typeof parent != 'number') {
			throw 'invalid type for parent';
		}
		const resp = await axios.get(`${API_URL}/getcomment?id=${parent}`);
		return resp.data;
	}
	async function sendReply(e: Event) {
		e.preventDefault();
		const resp = await axios.post(
			`${API_URL}/newcomment`,
			{ text: replyText, parent_id: parent },
			{ withCredentials: true }
		);
		resp.data.user_avatar = user.current.Avatar;
		resp.data.user_name = user.current.Name;
		loadedComments.push(resp.data);
	}

	async function getCommentsForParent(parent: number | null): Promise<any> {
		let parentStr: string;
		if (parent == null) {
			parentStr = 'null';
		} else if (typeof parent == 'number') {
			parentStr = `${parent}`;
		} else {
			throw 'invalid type for parent';
		}
		const resp = await axios.get(
			`${API_URL}/comments?parent_id=${parentStr}&pageSize=${PAGESIZE}&page=${page}`
		);
		if (!resp.data) {
			return null;
		}
		return resp.data.data;
	}

	async function loadMoreComments(initial: boolean = false) {
		if (!initial) {
			page++;
		} else {
			page = 1;
		}
		const comments = await getCommentsForParent(parent);

		if (comments == null || comments.length < PAGESIZE) {
      moreComments = false;
    }

		if (initial) {
			loadedComments = comments;
		} else {
			loadedComments = loadedComments.concat(comments);
		}
	}
</script>

<div class="mb-2 w-full">
	{#if parent == null}
		<h2 class="text-xl font-semibold">Comments</h2>
	{:else}
		<h2 class="mb-4 text-xl font-semibold">Following Thread</h2>
		{#await getThreadComment() then comment}
			<Comment {comment} />
		{/await}
	{/if}
</div>
<div class="w-full rounded-xl bg-gradient-to-r from-primary to-secondary p-0.5">
	<form
		class="flex max-h-6 min-h-7 w-full flex-col items-center overflow-y-hidden rounded-xl transition-all focus-within:min-h-45 {replyText.length >
		0
			? 'min-h-45'
			: null}"
		onsubmit={(e) => sendReply(e)}
	>
		<textarea
			class="min-h-5 w-full flex-1 bg-white px-2 py-0.5 transition-all outline-none placeholder:text-primary"
			placeholder="Write a comment"
			bind:value={replyText}
		></textarea>
		<div class="flex w-full justify-end gap-2 bg-white p-2" tabindex="-1">
			<button class="rounded-xl border p-2" tabindex="0" onclick={() => (replyText = '')}
				>Cancel</button
			>
			<button class="rounded-xl bg-secondary px-4" tabindex="0" type="submit">Comment</button>
		</div>
	</form>
</div>
<div class="flex w-full flex-col {className}">
	{#await loadMoreComments(true) then}
		<div class="flex w-full flex-col items-center gap-4 pt-8">
			{#each loadedComments as _, i}
				<Comment bind:comment={loadedComments[i]} />
			{/each}
			{#if moreComments}
				<Button class="max-w-32" onclick={() => loadMoreComments()}>Load More</Button>
			{/if}
		</div>
	{/await}
</div>
