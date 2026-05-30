import { publicApiBaseUrl } from "../config";

type ApiErrorPayload={ 
    message?:string
};

export async function clientApiFetch<T>(
    path: string,
    init?:RequestInit
):Promise<T>{
    const cleanPath = path.startsWith('/') ? path.slice(1) : path;
    const response=await fetch(`${publicApiBaseUrl}/${cleanPath}`,{
        ...init,
        credentials:'include',
        headers:{
            "content-Type":"application/json",
            ...(init?.headers ?? {})
        }
    })
    if(!response.ok){
        let payload:ApiErrorPayload| null=null;
        try{
            payload=(await response.json())as ApiErrorPayload;
        }catch(err){
            payload=null
        }
        let message = payload?.message ?? `Request failed with status ${response.status}`;
        throw new Error(message);
    }
    return (await response.json()) as T;

}