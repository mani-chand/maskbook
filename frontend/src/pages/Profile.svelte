<script lang="ts">
  import { onMount } from "svelte";
  import { Card } from "flowbite-svelte";

  // 1. Initialize state as null
  let LoggedInUser: any = $state(null);

  // 2. Fetch from LocalStorage ONLY when component mounts in the browser
  onMount(() => {
    const storedUser = localStorage.getItem("user");
    if (storedUser) {
      try {
        LoggedInUser = JSON.parse(storedUser);
      } catch (e) {
        console.error("Error parsing user data", e);
      }
    }
  });
</script>

<main class="flex justify-center p-6">
  {#if LoggedInUser}
    <Card
      class="w-full max-w-sm shadow-md border-gray-200 dark:border-gray-700 p-0 overflow-hidden"
    >
      <div class="h-24 bg-gradient-to-r from-purple-500 to-indigo-600"></div>

      <div class="flex flex-col items-center -mt-12 pb-6 px-4">
        <div
          class="relative w-24 h-24 rounded-full border-4 border-white dark:border-gray-800 shadow-sm bg-gray-100"
        >
          {#if LoggedInUser.avatar}
            <img
              src={LoggedInUser.avatar}
              alt={LoggedInUser.username}
              class="w-full h-full rounded-full object-cover"
            />
          {:else}
            <div
              class="w-full h-full rounded-full flex items-center justify-center text-3xl font-bold text-gray-400"
            >
              {LoggedInUser.username?.charAt(0).toUpperCase()}
            </div>
          {/if}
        </div>

        <div class="text-center mt-3">
          <h2 class="text-xl font-bold text-gray-900 dark:text-white">
            {LoggedInUser.username}
          </h2>
          <p class="text-sm text-gray-500 dark:text-gray-400 font-medium">
            {LoggedInUser.email}
          </p>
        </div>

        <div class="mt-4">
          <span
            class="px-3 py-1 text-xs font-semibold text-purple-700 bg-purple-100 rounded-full dark:bg-purple-900 dark:text-purple-300"
          >
            Active User
          </span>
        </div>
      </div>
    </Card>
  {:else}
    <div class="text-gray-500">Loading user profile...</div>
  {/if}
</main>
