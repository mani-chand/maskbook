import axios from 'axios';
// import type { NewPostForm } from './../store/store'; // You might need to adjust this type
import { fileToBase64 } from './../utils/file';

// Update input type to accept FileList (which is what Svelte binds return)
export const CreateNewPost = async (message: string, fileData: FileList | File[] | null) => {
    
    // 1. Initialize the object (This creates the base structure)
    // We use 'any' here temporarily because we are about to swap File[] for string[]
    const payload: any = {
        message: message,
        fileData: null, 
        user_id: JSON.parse(window.localStorage.getItem('user') || '{}').id,
    };

    if (fileData) {
        // 2. CONVERSION: Turn FileList into a standard Array
        const filesArray = Array.from(fileData);

        // 3. EXECUTION: Map over 'filesArray', NOT 'fileData'
        // This returns string[] (Base64 strings)
        const base64Files = await Promise.all(filesArray.map(fileToBase64));
        
        // 4. ASSIGNMENT: Assign the strings to the payload
        payload.fileData = base64Files;
    }

    axios.post('/api/createPost', payload, {
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
    })
    .then(res => { console.log('res', res); })
    .catch(err => { console.log(err); });
}