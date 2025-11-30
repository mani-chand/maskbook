<script>
  // @ts-nocheck

  import FileUpload from "../components/FileUpload.svelte";
  import { Label, Textarea, Button, Card, Avatar } from "flowbite-svelte";
  import { CreateNewPost } from "./../services/CreatePost";
  import { GetAllPost } from "./../services/GetAllPosts";
  import { onMount } from "svelte"; // Import onMount
  // State
  let posts = $state([]);
  let newPost = $state({
    message: "",
    fileData: null,
  });
  let LoggedInUser = $state(JSON.parse(localStorage.getItem("user")));

  let textareaprops = {
    id: "message",
    name: "message",
    rows: 3,
    placeholder: "What's on your mind?",
    class:
      "p-0 border-none focus:ring-0 bg-transparent text-lg w-full resize-none", // Clean look
  };

  const createPost = () => {
    CreateNewPost(newPost?.message, newPost?.fileData);
    // Optional: Reset after submit
    newPost.message = "";
    newPost.fileData = null;
  };

  // 2. Fetch data when component loads
  onMount(async () => {
    try {
      const data = await GetAllPost();
      // Ensure data is an array before assigning
      if (Array.isArray(data)) {
        posts = data;
        console.log($state.snapshot(posts), "posts");
      }
    } catch (error) {
      console.error("Failed to load posts:", error);
    }
  });
</script>

<main class="min-h-screen bg-gray-50 dark:bg-gray-900 py-8 px-4">
  <div class="max-w-2xl mx-auto space-y-6">
    <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-6">
      Social Feed
    </h1>

    <div
      class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 overflow-hidden"
    >
      <div
        class="p-4 border-b border-gray-100 dark:border-gray-700 bg-gray-50 dark:bg-gray-700/50"
      >
        <Label
          for="message"
          class="text-sm font-semibold text-gray-500 dark:text-gray-400 uppercase tracking-wider"
        >
          Create a Post
        </Label>
      </div>

      <div class="p-4">
        <div class="flex gap-4">
          <div class="flex-shrink-0">
            <div
              class="w-10 h-10 rounded-full bg-purple-100 text-purple-600 flex items-center justify-center font-bold"
            >
              <img
                src={LoggedInUser?.avatar}
                alt="Profile pic"
                class="w-full h-full rounded-full object-cover"
              />
            </div>
          </div>

          <div class="flex-grow">
            <Textarea {...textareaprops} bind:value={newPost.message} />
          </div>
        </div>
      </div>

      <div
        class="bg-gray-50 dark:bg-gray-700/30 p-3 flex items-center justify-between border-t border-gray-100 dark:border-gray-700"
      >
        <div class="flex items-center">
          <div class="scale-90 origin-left">
            <FileUpload multiple={true} bind:newPost />
          </div>
        </div>

        <Button
          onclick={createPost}
          color="purple"
          class="px-6 font-medium rounded-lg shadow-sm hover:shadow transition-all"
          disabled={!newPost?.fileData && !newPost?.message}
        >
          Post
        </Button>
      </div>
    </div>

    <div class="space-y-6">
      {#each posts as post}
        <Card
          class="w-full max-w-none shadow-sm border border-gray-200 dark:border-gray-700 overflow-hidden"
        >
          <div class="flex items-center gap-3 mb-3 px-2 pt-2">
            {#if post.author && post.author.avatar}
              <img
                src={post.author.avatar}
                alt={post.author.username}
                class="w-10 h-10 rounded-full object-cover ring-2 ring-gray-100 dark:ring-gray-700"
              />
            {/if}

            <div class="flex flex-col">
              <h3
                class="font-bold text-gray-900 dark:text-white text-base leading-tight"
              >
                {post.author ? post.author.username : "Unknown User"}
              </h3>
              <span
                class="text-xs text-gray-500 dark:text-gray-400 font-medium"
              >
                {post.created_at}
              </span>
            </div>
          </div>

          <div class="px-2 mb-3">
            <p
              class="text-gray-800 dark:text-gray-200 text-base leading-relaxed"
            >
              {post.message}
            </p>
          </div>

          {#if post.fileData && post.fileData.length > 0}
            {#each post.fileData as file}
              <div
                class="-mx-6 -mb-6 bg-gray-100 dark:bg-gray-800 border-t border-gray-100 dark:border-gray-700 mt-3"
              >
                <img
                  src={file}
                  alt="Post attachment"
                  class="w-full h-auto max-h-96 object-cover object-top hover:opacity-95 transition-opacity cursor-pointer"
                />
              </div>
            {/each}
          {/if}
        </Card>
      {/each}
    </div>
  </div>
</main>
