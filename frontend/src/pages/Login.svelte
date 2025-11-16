<script lang="ts">
  import { Input, Label, Button, Checkbox, A, Heading } from "flowbite-svelte";
</script>

<main
  style="width: 70%;height:70vh;display:flex;flex-direction:column;justify-content:center;align-content:center;margin:0 auto;"
>
  <form style="margin-top: 20px;">
    <div class="mb-6 grid gap-6 md:grid-cols-1">
      <div>
        <Label for="first_name" class="mb-2">Username</Label>
        <Input type="text" id="first_name" placeholder="John" required />
      </div>
      <div class="mb-6">
        <Label for="password" class="mb-2">Password</Label>
        <Input type="password" id="password" placeholder="•••••••••" required />
      </div>
      <Checkbox classes={{ div: "mb-6 gap-1 rtl:space-x-reverse" }} required>
        I agree with the <A
          href="/"
          class="text-primary-700 dark:text-primary-600 hover:underline"
          >terms and conditions</A
        >.
      </Checkbox>
      <Button type="submit">Signin</Button>
    </div>
  </form>
</main>
