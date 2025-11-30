import axios from 'axios'

// 1. Mark function as async
export const GetAllUsers = async () => {
    try {
        // 2. Wait for the response
        const res = await axios.get('/api/users', {
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
            },
        });
        
        console.log(res.data);
        // 3. Return the actual data
        return res.data; 

    } catch (err) {
        console.log('Somethings went wrong', err);
        return [];
    }
}