<script lang="ts">
	import { goto } from '$app/navigation';

	let first_name = '';
	let last_name = '';
	
    let email = '';
	let password = '';

	let registerPromise: Promise<any> | null = null

	function register() {
		registerPromise = fetch('api/register', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json'},
			credentials: 'include',
			body: JSON.stringify({ first_name, last_name, email, password })
		}).then(async (res) => {
			if (!res.ok) {
				const err = await res.json();
				throw new Error(err.message || 'Registration failed');
			}
			return res.json();
		}).then((data) => {

			setTimeout( () => {
				goto('/dashboard')
			}, 1500)

			return data;
      });
	}

</script>

<title>Register</title>

<form class="max-w-md mx-auto mt-10 p-6 bg-white rounded-xl shadow-lg space-y-4" on:submit={register}>

	<h2 class="text-2xl font-bold text-center">Register</h2>

	<div>
		<label for="first_name" class="block text-sm font-medium text-gray-700 mb-1">First name</label>
		<input
			id="first_name"
			type="text"
			bind:value={first_name}
			required
			class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-stone-500 focus:border-stone-500"
		/>
	</div>

	<div>
		<label for="last_name" class="block text-sm font-medium text-gray-700 mb-1">Last name</label>
		<input
			id="last_name"
			type="text"
			bind:value={last_name}
			required
			class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-stone-500 focus:border-stone-500"
		/>
	</div>

	<div>
		<label for="email" class="block text-sm font-medium text-gray-700 mb-1">Email</label>
		<input
			id="email"
			type="email"
			bind:value={email}
			required
			class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-stone-500 focus:border-stone-500"
		/>
	</div>

	<div>
		<label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
		<input
			id="password"
			type="password"
			bind:value={password}
			required
			class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-stone-500 focus:border-stone-500"
		/>
	</div>

	<button
		type="submit"
		class="w-full py-2 px-4 bg-stone-300 hover:bg-stone-400 text-white font-semibold rounded-md transition-colors"
	>
		Register
	</button>

	<p class="text-sm text-gray-500 text-center">
		Already have an account? <a href="/login" class="text-stone-600 hover:underline">Login</a>
	</p>

	{#if registerPromise}
	{#await registerPromise}
		<p>Registering...</p>
	{:then data}
		<p class="text-green-500">Success! Welcome {data.first_name}</p>
	{:catch error}
		<p class="capitalize text-red-500">{error.message}</p>
	{/await}
{/if}

</form>
