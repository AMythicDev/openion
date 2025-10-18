<script lang="ts">
	import { API_URL } from '$lib';
	let { class: className = null, parent = null } = $props();
	import axios from 'axios';
	import Comment from './Comment.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	let loadedComments: any = $state([]);
	let page = 0;
	let moreComments = $state(true);
	const PAGESIZE: number = 20;

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
		return resp.data.data;
	}

	async function loadMoreComments(initial: boolean = false) {
		page++;
		const comments = await getCommentsForParent(parent);
		if (comments == null) return (moreComments = false);
		if (initial) {
			loadedComments = loadedComments.concat(comments);
		}
	}
</script>

<div class="flex w-full flex-col {className}">
	{#await loadMoreComments(true) then}
		<div class="flex w-full flex-col items-center gap-4 pt-8">
			{#each loadedComments as comment}
				<Comment {comment} />
			{/each}
			{#if moreComments && loadedComments.length == page * PAGESIZE}
				<Button class="max-w-32" onclick={loadMoreComments}>Load More</Button>
			{/if}
		</div>
	{/await}
</div>
