<script lang="ts">
	let { comment = $bindable(), class: className = null } = $props();
	import * as Avatar from '$lib/components/ui/avatar/index.js';
	import { API_URL, shortenName } from '$lib';
	import { ThumbsUp, MessageCircle } from '@lucide/svelte';
	import { goto } from '$app/navigation';
	import axios from 'axios';
	import humanizeDuration from 'humanize-duration';

	const curtime = new Date();
	const comtime = new Date(comment.created_at);
	const dur = humanizeDuration(curtime - comtime, { largest: 1 });

	function goToComment() {
		goto(`?parent=${comment.id}`);
	}

	async function upvote() {
		await axios.patch(
			`${API_URL}/updatepost`,
			{ id: comment.id, upvotes: comment.upvotes + 1 },
			{ withCredentials: true }
		);
		comment.upvotes += 1;
	}
</script>

<section class="w-full rounded-xl border border-primary {className}">
	<div class="flex items-center gap-2 rounded-t-xl bg-primary/10 p-2">
		<Avatar.Root class="h-8 w-8 shadow-sm">
			<Avatar.Image src={comment.user_avatar} alt="user image" />
			<Avatar.Fallback>{shortenName(comment.user_name)}</Avatar.Fallback>
		</Avatar.Root>
		<span class="font-bold text-primary">{comment.user_name}</span>
		<span class="translate-y-0.5 rounded-b-xl text-xs text-primary/90">{dur} ago</span>
	</div>
	<p class="p-2">
		{comment.text}
	</p>
	<div class="mt-2 flex gap-3 px-2 pb-2">
		<button
			class="flex max-w-max items-center gap-2 rounded-xl border-2 border-gray-500 px-2.5 py-0.5"
			onclick={upvote}
		>
			<ThumbsUp size="14" class="stroke-gray-500" />
			{comment.upvotes}
		</button>
		<button
			class="flex max-w-max items-center gap-2 rounded-xl border-2 border-gray-500 px-2.5 py-0.5"
			onclick={goToComment}
		>
			<MessageCircle size="14" class="stroke-gray-500" />
			{comment.replies}
		</button>
	</div>
</section>
