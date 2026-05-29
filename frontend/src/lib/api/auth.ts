import {AuthMeResponse, User,UserRole } from "@/types/auth"
// import { serverApiBaseUrl } from "../config";
import { serverApifetch } from "./server";
import { redirect } from "next/navigation";

export async function getUserInfo():Promise<User | null>{
  try{
   const auth =(await serverApifetch<AuthMeResponse>("/auth/me"))
return auth.user ?? null;
  }
  catch(error){
    console.error("Faliure to fetch user info: ",error);
    return null;
  }
};

export async function requireUser(allowedRoles?:UserRole[]):Promise<User>{
    const user=await getUserInfo();
    if(!user){
        redirect("/")
    }
    if(allowedRoles &&!allowedRoles.includes(user.role)){
        redirect (user.role==="admin"?"/admin":"/submission");
    }
    return user;
}