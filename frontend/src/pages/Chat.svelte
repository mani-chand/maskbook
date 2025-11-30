<script>
  import UsersList from "../components/UsersList.svelte";
  import ChatWindow from "../components/ChatWindow.svelte";
  // import Topbar from "../components/Topbar.svelte"; // Assuming you have this

  let currentUser = $state(null);

  const handleUserSelection = (user) => {
    currentUser = user;
  };
</script>

<main class="h-[calc(90vh)] w-full overflow-hidden flex">
  <aside
    class="w-80 h-full flex-shrink-0 border-r border-gray-200 dark:border-gray-700"
  >
    <UsersList onUserSelect={handleUserSelection} />
  </aside>

  <section class="flex-1 h-full">
    <ChatWindow selectedUser={currentUser} />
  </section>
</main>
