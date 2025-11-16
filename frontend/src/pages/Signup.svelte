<script lang="ts">
  import { Input, Label, Button, Checkbox, A, Heading } from "flowbite-svelte";
</script>

<main style="width: 70%;display:flex;flex-direction:column;margin:0 auto;">
  <form style="margin-top: 20px;">
    <div class="mb-6 grid gap-6 md:grid-cols-2">
      <div>
        <Label for="first_name" class="mb-2">Username</Label>
        <Input type="text" id="first_name" placeholder="John" required />
      </div>
      <div>
        <Label for="email" class="mb-2">Email address</Label>
        <Input
          type="email"
          id="email"
          placeholder="john.doe@company.com"
          required
        />
      </div>
    </div>
    <div class="mb-6">
      <Label for="phone" class="mb-2">Phone number</Label>
      <Input
        type="tel"
        id="phone"
        placeholder="123-45-678"
        pattern={"[0-9]{3}-[0-9]{2}-[0-9]{3}"}
        required
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
      />
    </div>
    <Checkbox classes={{ div: "mb-6 gap-1 rtl:space-x-reverse" }} required>
      I agree with the <A
        href="/"
        class="text-primary-700 dark:text-primary-600 hover:underline"
        >terms and conditions</A
      >.
    </Checkbox>
    <Button type="submit">Signup</Button>
  </form>
</main>
