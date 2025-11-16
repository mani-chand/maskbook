import type { NewPostForm } from './../store/store'
export const CreateNewPost = (message:string,fileData: FileList | null)=>{
    var newPost:NewPostForm = {
        message:message,
        fileData:fileData,
    }

    
} 