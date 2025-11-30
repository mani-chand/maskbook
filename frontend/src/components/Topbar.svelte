<script lang="ts">
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger,
  } from "flowbite-svelte";
  let LoggedInUser = JSON.parse(localStorage.getItem("user"));
</script>

<Navbar
  class="sticky start-0 top-0 z-20 w-full bg-white px-2 py-2.5 sm:px-4 dark:bg-gray-700"
>
  <NavBrand href="/">
    <span
      class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
      >Maskbook</span
    >
  </NavBrand>
  <NavHamburger />
  <NavUl>
    {#if !LoggedInUser}
      <NavLi href="/login">Login</NavLi>
      <NavLi href="/create">create an account</NavLi>
    {:else}
      <NavLi href="">Logout</NavLi>
      <NavLi href="/chat">chat</NavLi>
      <NavLi href="/profile/{LoggedInUser?.id}/">
        <img
          src={LoggedInUser?.avatar}
          alt={LoggedInUser?.username}
          class="w-10 h-10 rounded-full object-cover ring-2 ring-gray-100 dark:ring-gray-700"
        />
      </NavLi>
    {/if}
  </NavUl>
</Navbar>
