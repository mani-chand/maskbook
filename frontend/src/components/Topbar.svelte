<script lang="ts">
  import { onMount } from "svelte";
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger,
  } from "flowbite-svelte";

  let LoggedInUser = $state(null);

  onMount(() => {
    const user = localStorage.getItem("user");
    if (user) LoggedInUser = JSON.parse(user);
  });
</script>

<Navbar
  fluid
  class="start-0 top-0 z-20 w-full bg-white px-2 py-1 h-12 dark:bg-gray-700 border-b border-gray-200 dark:border-gray-600 flex items-center"
>
  <NavBrand href="/" class="flex items-center">
    <span
      class="self-center text-base font-bold whitespace-nowrap dark:text-white"
    >
      Maskbook
    </span>
  </NavBrand>

  <NavHamburger class="h-8 w-8" />

  <NavUl class="flex items-center">
    {#if !LoggedInUser}
      <NavLi href="/login" class="py-1">Login</NavLi>
      <NavLi href="/create" class="py-1">Create</NavLi>
    {:else}
      <NavLi href="/logout" class="py-1">Logout</NavLi>
      <NavLi href="/chat" class="py-1">Chat</NavLi>

      <NavLi
        href="/profile/{LoggedInUser?.id}/"
        class="p-0 flex items-center ml-2"
      >
        {#if LoggedInUser?.avatar}
          <img
            src={LoggedInUser.avatar}
            alt="User"
            class="w-7 h-7 rounded-full object-cover ring-1 ring-gray-200 dark:ring-gray-600"
          />
        {:else}
          <div
            class="w-7 h-7 rounded-full bg-purple-100 flex items-center justify-center text-[10px] font-bold text-purple-600"
          >
            {LoggedInUser?.username?.charAt(0).toUpperCase()}
          </div>
        {/if}
      </NavLi>
    {/if}
  </NavUl>
</Navbar>
