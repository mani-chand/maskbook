<script>
  let { selectedUser } = $props();
</script>

{#if selectedUser}
  <div class="flex flex-col h-full bg-gray-50 dark:bg-gray-900">
    <div
      class="p-4 bg-white dark:bg-gray-800 border-b border-gray-200 dark:border-gray-700 flex items-center gap-3 shadow-sm"
    >
      {#if selectedUser.avatar}
        <img
          src={selectedUser.avatar}
          alt="avatar"
          class="w-10 h-10 rounded-full object-cover"
        />
      {:else}
        <div
          class="w-10 h-10 rounded-full bg-purple-100 flex items-center justify-center text-purple-600 font-bold text-sm"
        >
          {selectedUser?.username?.slice(0, 2).toUpperCase()}
        </div>
      {/if}
      <h2 class="text-lg font-bold text-gray-800 dark:text-white">
        {selectedUser.username}
      </h2>
    </div>

    <div
      class="flex-1 p-6 overflow-y-auto flex items-center justify-center text-gray-400"
    >
      <p>This is the start of your conversation with {selectedUser.username}</p>
    </div>

    <div
      class="p-4 bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700"
    >
      <input
        type="text"
        placeholder="Type a message..."
        class="w-full p-2 border rounded-lg dark:bg-gray-700 dark:text-white focus:ring-purple-500 focus:border-purple-500"
      />
    </div>
  </div>
{:else}
  <div
    class="flex items-center justify-center h-full bg-gray-50 dark:bg-gray-900 text-gray-500"
  >
    <p>Select a user to start chatting</p>
  </div>
{/if}
