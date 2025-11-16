import { writable, derived } from "svelte/store";

export const posts = writable([]);
export const newPost = writable({});
export const user = writable({});