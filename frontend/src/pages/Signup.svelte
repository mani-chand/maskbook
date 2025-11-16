<script lang="ts">
  import { Input, Label, Button, Checkbox, A } from "flowbite-svelte";
  import FileUpload from "../components/FileUpload.svelte";
  import { CreateUser } from "./../services/createUser";
  let newUser = $state({
    username: "",
    email: "",
    mobile: "",
    password: "",
    confirmPassword: "",
    fileData: null,
  });
</script>

<main style="width: 70%;display:flex;flex-direction:column;margin:0 auto;">
  <form style="margin-top: 20px;">
    <div class="mb-6 grid gap-6 md:grid-cols-2">
      <div>
        <Label for="username" class="mb-2">Username</Label>
        <Input
          bind:value={newUser.username}
          type="text"
          id="username"
          placeholder="John"
          required
        />
      </div>
      <div>
        <Label for="email" class="mb-2">Email address</Label>
        <Input
          type="email"
          id="email"
          placeholder="john.doe@company.com"
          required
          bind:value={newUser.email}
        />
      </div>
    </div>
    <div class="mb-6">
      <Label for="phone" class="mb-2">Phone number</Label>
      <Input
        type="tel"
        id="phone"
        placeholder="+1 12345678"
        required
        bind:value={newUser.password}
      />
    </div>
    <div class="mb-6">
      <Label for="password" class="mb-2">Password</Label>
      <Input type="password" id="password" placeholder="•••••••••" required />
    </div>
    <div class="mb-6">
      <Label for="confirm_password" class="mb-2">Confirm password</Label>
      <Input
        type="password"
        id="confirm_password"
        placeholder="•••••••••"
        required
        bind:value={newUser.confirmPassword}
      />
    </div>
    <div class="mb-6">
      <Label>Profile Picture</Label>
      <FileUpload multiple={false} bind:newPost={newUser} />
    </div>
    <Checkbox classes={{ div: "mb-6 gap-1 rtl:space-x-reverse" }} required>
      I agree with the <A
        href="/"
        class="text-primary-700 dark:text-primary-600 hover:underline"
        >terms and conditions</A
      >.
    </Checkbox>
    <Button
      onclick={() => {
        CreateUser(newUser);
      }}
      type="submit">Signup</Button
    >
  </form>
</main>
