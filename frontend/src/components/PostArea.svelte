<script>
  import FileUpload from "../components/FileUpload.svelte";
  import { Label, Textarea } from "flowbite-svelte";
  let textareaprops = {
    id: "message",
    name: "message",
    label: "Your message",
    rows: 3,
    placeholder: "Write something...",
  };
</script>

<main>
  <div style="width: 80%;display:flex;flex-direction:column;margin:0 auto;">
    <Label style="font-size: large;">Create a Post</Label>
    <Textarea {...textareaprops} class="w-full" />
    <FileUpload />
  </div>
</main>
