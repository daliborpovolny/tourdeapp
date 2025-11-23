<script lang="ts">
	import { page } from '$app/state';

	let coursesPromise: Promise<any> = loadCourseDetail();

	async function loadCourseDetail() {
		return fetch('/api/courses/' + page.params.uuid, {
			method: 'GET',
			headers: { 'Content-type': 'application/json' }
		})
			.then(async (res) => {
				if (res.status == 404) {
					throw new Error('Unknown course');
				}

				if (!res.ok) {
					try {
						const err = await res.json();
					} catch {
						throw new Error('Failed to get course info');
					}
				}
				return res.json();
			})
			.then((data) => {
				return data;
			});
	}
</script>

<br />
{#await coursesPromise}
	<p>Loading course detail...</p>
{:then data}
	<div>
		<h2>{data.name}</h2>
		<br />
		<p>{data.description}</p>
	</div>
{:catch error}
	<p class="text-red-500 capitalize">{error.message}</p>
{/await}
