<script lang="ts">
  import { Fileupload, label } from "flowbite-svelte";

  let { newPost = $bindable() } = $props();
  let fileuploadprops = {
    id: "user_avatar",
    label: "Upload file",
  };
</script>

<Fileupload bind:files={newPost.file} {...fileuploadprops} />
