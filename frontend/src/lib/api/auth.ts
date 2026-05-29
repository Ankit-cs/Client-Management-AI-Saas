import {AuthMeResponse, User,UserRole } from "@/types/auth"
// import { serverApiBaseUrl } from "../config";
import { serverApifetch } from "./server";
import { redirect } from "next/navigation";


function isUnauthorized(error: unknown):boolean{
  if (!(error instanceof Error)) return false;
  const msg = error.message.toLowerCase();
  return msg.includes('unauthorized') || msg.includes('unauthorised') || msg.includes('fetch failed') || msg.includes('econnrefused');
}


export async function getUserInfo():Promise<User | null>{
  try{
   const auth =(await serverApifetch<AuthMeResponse>("/auth/me"))
return auth.user ?? null;
  }
  catch(error){
   if(isUnauthorized(error)){
    // console.log("No authenticated user");
    return null;
   }
   throw error;
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