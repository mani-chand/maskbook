<script lang="ts">
  import { GetAllUsers } from "./../services/GetAllUsers";
  import { onMount } from "svelte";

  let { onUserSelect } = $props();
  let users = $state([]);

  onMount(async () => {
    try {
      const data = await GetAllUsers();
      if (Array.isArray(data)) users = data;
    } catch (error) {
      console.error("Failed to load users:", error);
    }
  });
</script>

<div class="flex flex-col h-full bg-white dark:bg-gray-800">
  <div class="p-4 border-b border-gray-200 dark:border-gray-700 flex-shrink-0">
    <h2 class="text-xl font-bold text-gray-800 dark:text-white">Users</h2>
  </div>

  <div class="flex-1 overflow-y-auto custom-scrollbar">
    {#each users as user}
      <button
        onclick={() => onUserSelect(user)}
        class="w-full flex items-center gap-3 p-3 hover:bg-purple-50 dark:hover:bg-gray-700 transition-colors cursor-pointer text-left focus:outline-none focus:bg-purple-100 border-b border-gray-50 dark:border-gray-700/50 last:border-0"
      >
        <div class="relative w-10 h-10 flex-shrink-0">
          {#if user?.avatar}
            <img
              src={user.avatar}
              alt={user.username}
              class="w-full h-full rounded-full object-cover border border-gray-200"
            />
          {:else}
            <div
              class="w-full h-full rounded-full bg-purple-100 flex items-center justify-center text-purple-600 font-bold text-sm"
            >
              {user?.username?.slice(0, 2).toUpperCase()}
            </div>
          {/if}
          <span
            class="bottom-0 left-7 absolute w-3 h-3 bg-green-400 border-2 border-white dark:border-gray-800 rounded-full"
          ></span>
        </div>

        <div class="min-w-0">
          <p
            class="text-sm font-semibold text-gray-900 dark:text-white truncate"
          >
            {user?.username}
          </p>
          <p class="text-xs text-gray-500 dark:text-gray-400 truncate">
            Click to start chatting
          </p>
        </div>
      </button>
    {/each}
  </div>
</div>

<style>
  /* Modern thin scrollbar styling */
  .custom-scrollbar::-webkit-scrollbar {
    width: 6px;
  }
  .custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
  }
  .custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: #d1d5db; /* gray-300 */
    border-radius: 20px;
  }
  /* Dark mode scrollbar support */
  :global(.dark) .custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: #4b5563; /* gray-600 */
  }
</style>
