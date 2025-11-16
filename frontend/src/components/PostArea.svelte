<script>
  import FileUpload from "../components/FileUpload.svelte";
  import { Label, Textarea } from "flowbite-svelte";
  import { Button } from "flowbite-svelte";
  let textareaprops = {
    id: "message",
    name: "message",
    label: "Your message",
    rows: 3,
    placeholder: "Write something...",
  };
  let newPost = $state({
    message: "",
  });
  const createPost = () => {};
</script>

<main>
  <div style="width: 80%;display:flex;flex-direction:column;margin:0 auto;">
    <Label style="font-size: large;">Create a Post</Label>
    <Textarea {...textareaprops} bind:value={newPost.message} class="w-full" />
    <FileUpload bind:newPost />
    <Button
      class="cursor-pointer"
      onclick={createPost}
      style="margin-top:3px;"
      color="purple">Submit</Button
    >
  </div>
</main>
