import axios from 'axios'
import type { NewPostForm } from './../store/store'
import {fileToBase64} from './../utils/file'
export const CreateNewPost = async (message:string,fileData: File[] | null)=>{
    var newPost:NewPostForm = {
        message: message,
        fileData: null,
        user_id: JSON.parse(window.localStorage.getItem('user')).id,
    }
    if (fileData) {
        const files = [...fileData]; // FileList → File[]
        newPost.fileData = await Promise.all(fileData.map(fileToBase64));
    }
    axios.post('/api/createPost',newPost,{
        headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    }).then(res=>{console.log('res',res)}).catch(err=>{console.log(err)});
} 