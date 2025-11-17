import axios from 'axios'
import type {loginForm} from './../store/store' 
import {user} from './../store/store'
export const ValidateUser = async (newUser: loginForm) => {
    axios.post('/api/login',newUser,{
        headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    }).then(res=>{
        user.set(res?.data?.user)
        console.log(user,"user logged in")
    }).catch(err=>{console.log(err,'error')})
}