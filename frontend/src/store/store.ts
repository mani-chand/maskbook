import { writable, type Writable } from 'svelte/store';

// --- Interface Definitions ---

/**
 * Represents a User object.
 */

export interface loginForm {
  username:string,
  password:string
}

export interface User {
  id: string;
  username: string;
  avatar?: string; // Optional avatar
  email:string;
  mobile:string;
  // Add other user fields as needed
}
export interface SignupForm {
  username: string;
  email: string;
  mobile: string;
  password: string;
  confirmPassword: string;
  // This is the correct type for "null or a single file"
  fileData: File | null;
}

/**
 * Represents a Post object as it's stored in the database.
 */
export interface Post {
  id: string;
  message: string;
  fileData: FileList | null; // Optional file URL
  user_id: string;
  created_at: string; // Typically an ISO string
  user?: User; // Optional: Include full user details
}

/**
 * Represents the data in the "Create a Post" form.
 */
export interface NewPostForm {
  message: string;
  fileData: FileList | null; // The FileList from the <Fileupload> component
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
  file: FileList | null,
});

/**
 * Stores the currently authenticated user.
 * It's conventional to use `null` if no user is logged in.
 */
export const user: Writable<User | null> = writable(null);