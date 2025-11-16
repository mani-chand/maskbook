import axios from 'axios'
import type {SignupForm} from './../store/store'
export const CreateUser = (newUser:SignupForm)=>{
    axios.post('/api/createUser',{
        data:newUser,
        headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        },
    }).then(res=>{
        console.log(res)
    }).catch(err=>{
        console.log(err)
    })
}