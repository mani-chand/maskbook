<script lang="ts">
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger,
  } from "flowbite-svelte";
</script>

<Navbar
  class="sticky start-0 top-0 z-20 w-full bg-white px-2 py-2.5 sm:px-4 dark:bg-gray-800"
>
  <NavBrand href="/">
    <span
      class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
      >Maskbook</span
    >
  </NavBrand>
  <NavHamburger />
  <NavUl>
    <NavLi href="/login">Login</NavLi>
    <NavLi href="">Logout</NavLi>
    <NavLi href="/create">create an account</NavLi>
    <NavLi href="/chat">chat</NavLi>
    <NavLi href="/profile/1/">profile</NavLi>
  </NavUl>
</Navbar>
