<script>
  // @ts-nocheck
  import FileUpload from "../components/FileUpload.svelte";
  import { Label, Textarea, Button, Card } from "flowbite-svelte";
  import { CreateNewPost } from "./../services/CreatePost";
  import { GetAllPost } from "./../services/GetAllPosts";
  import { onMount } from "svelte";

  let posts = $state([]);
  let newPost = $state({ message: "", fileData: null });
  let LoggedInUser = $state(null);

  // 1. New State for the Lightbox
  let selectedImage = $state(null);

  let textareaprops = {
    id: "message",
    name: "message",
    rows: 3,
    placeholder: "What's on your mind?",
    class:
      "p-0 border-none focus:ring-0 bg-transparent text-lg w-full resize-none",
  };

  const createPost = async () => {
    await CreateNewPost(newPost?.message, newPost?.fileData);
    newPost.message = "";
    newPost.fileData = null;
    const data = await GetAllPost();
    if (Array.isArray(data)) posts = data;
  };

  onMount(async () => {
    const user = localStorage.getItem("user");
    if (user) LoggedInUser = JSON.parse(user);

    try {
      const data = await GetAllPost();
      if (Array.isArray(data)) posts = data;
    } catch (error) {
      console.error("Failed to load posts:", error);
    }
  });

  // 2. Helper functions to open/close modal
  const openImage = (url) => {
    selectedImage = url;
    document.body.style.overflow = "hidden"; // Disable background scrolling
  };

  const closeImage = () => {
    selectedImage = null;
    document.body.style.overflow = "auto"; // Re-enable scrolling
  };

  // Close on Escape key
  const handleKeydown = (e) => {
    if (e.key === "Escape") closeImage();
  };
</script>

<svelte:window onkeydown={handleKeydown} />

<main class="h-[calc(90vh)] w-full flex flex-col bg-gray-50 dark:bg-gray-900">
  <div class="flex-1 overflow-y-auto px-4 pb-8 custom-scrollbar">
    <div class="max-w-2xl mx-auto space-y-6">
      <div
        class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 overflow-hidden mt-4"
      >
        <div
          class="p-4 border-b border-gray-100 dark:border-gray-700 bg-gray-50 dark:bg-gray-700/50"
        >
          <Label
            class="text-sm font-semibold text-gray-500 dark:text-gray-400 uppercase tracking-wider"
          >
            Create a Post
          </Label>
        </div>
        <div class="p-4">
          <div class="flex gap-4">
            <div class="flex-shrink-0">
              <div
                class="w-10 h-10 rounded-full bg-purple-100 text-purple-600 flex items-center justify-center font-bold overflow-hidden"
              >
                {#if LoggedInUser?.avatar}
                  <img
                    src={LoggedInUser.avatar}
                    alt="Profile"
                    class="w-full h-full object-cover"
                  />
                {:else}
                  <span>ME</span>
                {/if}
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
              <div
                class="-mx-6 -mb-6 mt-3 border-t border-gray-100 dark:border-gray-700 bg-gray-100 dark:bg-gray-800"
              >
                <div
                  class="grid gap-0.5 overflow-hidden"
                  class:grid-cols-1={post.fileData.length === 1}
                  class:grid-cols-2={post.fileData.length >= 2}
                  class:h-96={post.fileData.length > 1}
                >
                  {#each post.fileData.slice(0, 4) as file, i}
                    <button
                      class="relative w-full h-full overflow-hidden focus:outline-none"
                      onclick={() => openImage(file)}
                    >
                      <img
                        src={file}
                        alt="Post attachment"
                        class="w-full h-full object-cover hover:opacity-90 transition-opacity cursor-pointer"
                        style={post.fileData.length === 1
                          ? "max-height: 24rem;"
                          : "height: 100%;"}
                      />

                      {#if i === 3 && post.fileData.length > 4}
                        <div
                          class="absolute inset-0 bg-black/60 flex items-center justify-center text-white font-bold text-2xl"
                        >
                          +{post.fileData.length - 4}
                        </div>
                      {/if}
                    </button>
                  {/each}
                </div>
              </div>
            {/if}
          </Card>
        {/each}
      </div>
    </div>
  </div>

  {#if selectedImage}
    <div
      class="fixed inset-0 z-[100] bg-black/95 backdrop-blur-sm flex items-center justify-center p-4"
      onclick={closeImage}
      role="button"
      tabindex="0"
      onkeypress={handleKeydown}
    >
      <button
        onclick={closeImage}
        class="absolute top-4 right-4 text-white/70 hover:text-white p-2 rounded-full hover:bg-white/10 transition-colors"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-8 w-8"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </button>

      <img
        src={selectedImage}
        alt="Full screen view"
        class="max-w-full max-h-[90vh] object-contain rounded-lg shadow-2xl"
        onclick={(e) => e.stopImmediatePropagation()}
      />
    </div>
  {/if}
</main>

<style>
  .custom-scrollbar::-webkit-scrollbar {
    width: 6px;
  }
  .custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
  }
  .custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: #d1d5db;
    border-radius: 20px;
  }
  :global(.dark) .custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: #4b5563;
  }
</style>
