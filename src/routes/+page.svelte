<script lang="ts">
	import CommentBox from './CommentBox.svelte';
  let { data } = $props();
  console.log(data.parent)
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
		<div class="mb-2 w-full">
			<h2 class="text-xl font-semibold">Comments</h2>
		</div>
		<div class="w-full rounded-xl bg-gradient-to-r from-primary to-secondary p-0.5">
			<form
				class="flex max-h-6 min-h-7 w-full flex-col items-center overflow-y-hidden rounded-xl transition-all focus-within:min-h-45"
			>
				<textarea
					class="min-h-5 w-full flex-1 bg-white px-2 py-0.5 transition-all outline-none placeholder:text-primary"
					placeholder="Write a comment"
				></textarea>
				<div class="flex w-full justify-end gap-2 bg-white p-2" tabindex="-1">
					<button class="rounded-xl border p-2" tabindex="0">Cancel</button>
					<button class="rounded-xl bg-secondary px-4" tabindex="0">Comment</button>
				</div>
			</form>
		</div>
    <CommentBox parent={data.parent} />
	</main>
</div>
