import axios from 'axios';
import type { SignupForm } from './../store/store';
import {fileToBase64} from './../utils/file'
// 2. Define the payload for the API
interface ApiPayload {
  username: string;
  email: string;
  mobile: string;
  password: string;
  fileData: string; // File is sent as a Base64 string
}

// 4. The CreateUser function
export const CreateUser = async (newUser: SignupForm) => {
  // Create a 'payload' object to send
  const payload: ApiPayload = {
    username: newUser.username,
    email: newUser.email,
    mobile: newUser.mobile,
    password: newUser.password,
    fileData: '', // Default to empty string
  };

  // If a file exists, read it as Base64
  if (newUser.fileData) {
    try {
      payload.fileData = await fileToBase64(newUser.fileData[0]);
    } catch (err) {
      console.error('Error reading file:', err);
      // Handle file read error (e.g., show to user)
      return;
    }
  }

  // Now send the 'payload' object
  try {
    const res = await axios.post('/api/createUser', payload, {
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    });
    console.log('User creation successful:', res.data);
    window.location.assign('/login')
    // You could add a redirect here on success
    // e.g., navigate('/login');
  } catch (err: any) {
    console.error('User creation failed:');
    if (err.response) {
      // The request was made and the server responded with a status code
      // that falls out of the range of 2xx
      console.error(err.response.data);
    } else if (err.request) {
      // The request was made but no response was received
      console.error(err.request);
    } else {
      // Something happened in setting up the request that triggered an Error
      console.error('Error', err.message);
    }
  }
};