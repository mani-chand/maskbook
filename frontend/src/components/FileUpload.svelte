<script lang="ts">
  import { Fileupload } from "flowbite-svelte";

  let fileuploadprops = {
    id: "user_avatar",
  };
</script>

<Fileupload {...fileuploadprops} />
