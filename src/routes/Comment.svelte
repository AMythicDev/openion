<script lang="ts">
	let { comment, class: className = null } = $props();
	import * as Avatar from '$lib/components/ui/avatar/index.js';
	import { shortenName } from '$lib';
	import { ThumbsUp, MessageCircle } from '@lucide/svelte';
	import { goto } from '$app/navigation';

	function goToComment() {
		throw goto(`?parent=${comment.id}`, { replaceState: true, invalidateAll: true });
	}
</script>

<section class="w-full rounded-xl border border-gray-500 {className}">
	<div class="flex items-center gap-2 rounded-t-xl bg-gray-300 p-2">
		<Avatar.Root class="h-8 w-8 shadow-sm">
			<Avatar.Image src={comment.user_avatar} alt="user image" />
			<Avatar.Fallback>{shortenName(comment.user_name)}</Avatar.Fallback>
		</Avatar.Root>
		{comment.user_name}
		<span class="translate-y-0.5 rounded-b-xl text-xs text-gray-400">{comment.created_at}</span>
	</div>
	<p class="p-2">
		{comment.text}
	</p>
	<div class="mt-2 flex gap-3 px-2 pb-2">
		<button
			class="flex max-w-max items-center gap-2 rounded-xl border-2 border-gray-500 px-2.5 py-0.5"
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
