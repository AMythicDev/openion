<script lang="ts">
	import CommentBox from './CommentBox.svelte';
	import type { PageProps } from './$types.js';
	const { data }: PageProps = $props();
</script>

<div class="flex flex-col items-center">
	<main
		class="mb-4 flex max-w-full flex-col items-center rounded-md bg-white p-8 shadow-sm md:max-w-1/2"
	>
		<h1 class="mb-4 w-full text-2xl font-bold">
			Lorem ipsum, dolor sit amet consectetur adipisicing elit.
		</h1>
		<div class="mb-4 leading-relaxed">
			<p class="mb-1.5">
				Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae
				pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu
				aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas.
				Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class
				aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.
			</p>
			<p>
				Lorem ipsum dolor sit amet consectetur adipiscing elit. Quisque faucibus ex sapien vitae
				pellentesque sem placerat. In id cursus mi pretium tellus duis convallis. Tempus leo eu
				aenean sed diam urna tempor. Pulvinar vivamus fringilla lacus nec metus bibendum egestas.
				Iaculis massa nisl malesuada lacinia integer nunc posuere. Ut hendrerit semper vel class
				aptent taciti sociosqu. Ad litora torquent per conubia nostra inceptos himenaeos.
			</p>
		</div>
		<CommentBox parent={data.parent} />
	</main>
</div>
