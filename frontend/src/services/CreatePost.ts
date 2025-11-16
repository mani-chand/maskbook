import type { NewPostForm } from './../store/store'
export const CreateNewPost = (message:string,file: FileList | null)=>{
    var newPost:NewPostForm = {
        message:message,
        file:file,
    }

    
} 