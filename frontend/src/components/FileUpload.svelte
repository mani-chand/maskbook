<script lang="ts">
  import { Fileupload } from "flowbite-svelte";

  let { newPost = $bindable(), multiple } = $props();
  let fileuploadprops = {
    id: "user_avatar",
    label: "Upload file",
    multiple: multiple,
  };
</script>

<Fileupload bind:files={newPost.fileData} {...fileuploadprops} />
