<script lang="ts">
	import { page } from '$app/state';

	let coursesPromise: Promise<any> = loadCourseDetail();

	async function loadCourseDetail() {
		return fetch('/api/courses/'+page.params.uuid, {
			method: 'GET',
			headers: { 'Content-type': 'application/json' }
		})
			.then(async (res) => {
				if (!res.ok) {
					const err = await res.json();
					throw new Error(err.message || 'Login failed');
				}
				return res.json();
			})
			.then((data) => {
				return data;
			});
	}
</script>

<br>
{#await coursesPromise}
	<p>Loading course detail...</p>
{:then data}
    <div>
        <h2>{data.name}</h2>
        <br>
        <p>{data.description}</p>
    </div>
{:catch error}
	<p></p>
{/await}
