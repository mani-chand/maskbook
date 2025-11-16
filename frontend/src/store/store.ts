import { writable, type Writable } from 'svelte/store';

// --- Interface Definitions ---

/**
 * Represents a User object.
 */
export interface User {
  id: string;
  username: string;
  avatar_url?: string; // Optional avatar
  // Add other user fields as needed
}

/**
 * Represents a Post object as it's stored in the database.
 */
export interface Post {
  id: string;
  message: string;
  file_url?: string; // Optional file URL
  user_id: string;
  created_at: string; // Typically an ISO string
  user?: User; // Optional: Include full user details
}

/**
 * Represents the data in the "Create a Post" form.
 */
export interface NewPostForm {
  message: string;
  file: FileList | null; // The FileList from the <Fileupload> component
}

// --- Store Definitions ---

/**
 * Stores an array of all posts.
 */
export const posts: Writable<Post[]> = writable([]);

/**
 * Stores the state of the "Create a Post" form.
 * We initialize it with default empty values.
 */
export const newPost: Writable<NewPostForm> = writable({
  message: "",
  file: null,
});

/**
 * Stores the currently authenticated user.
 * It's conventional to use `null` if no user is logged in.
 */
export const user: Writable<User | null> = writable(null);