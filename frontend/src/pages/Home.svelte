<script>
  import PostArea from "../components/PostArea.svelte";

  $effect(() => {
    let LoggedInUser = localStorage.getItem("user");
    if (!LoggedInUser) {
      window.location.assign("/login");
    }
  });
</script>

<main>
  <PostArea />
</main>
