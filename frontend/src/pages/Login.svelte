<script lang="ts">
  import { Input, Label, Button, Checkbox, A } from "flowbite-svelte";
  let userLogin = $state({
    username: "",
    password: "",
  });
</script>

<main
  style="width: 70%;height:70vh;display:flex;flex-direction:column;justify-content:center;align-content:center;margin:0 auto;"
>
  <form style="margin-top: 20px;">
    <div class="mb-6 grid gap-6 md:grid-cols-1">
      <div>
        <Label for="username" class="mb-2">Username</Label>
        <Input
          bind:value={userLogin.username}
          type="text"
          id="username"
          placeholder="John"
          required
        />
      </div>
      <div class="mb-6">
        <Label for="password" class="mb-2">Password</Label>
        <Input
          bind:value={userLogin.password}
          type="password"
          id="password"
          placeholder="•••••••••"
          required
        />
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
