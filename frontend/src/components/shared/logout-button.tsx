"use client";

import { useState } from "react";
import { Button, buttonVariants } from "../ui/button";
import { clientApiFetch } from "@/lib/api/client";
import { useRouter } from "next/navigation";
import { cn } from "@/lib/utils";

const logoutButtonClass = cn(
  "group cursor-pointer relative inline-flex items-center justify-center px-5 text-sm font-medium text-white bg-transparent border-2 border-[#414141] hover:scale-[1.03] hover:border-white transition-all duration-200 h-10 rounded-xl font-inter",
);

function LogoutButton() {
  const [isPending, setIsPending] = useState(false);
  const router = useRouter();

  async function handleLogout() {
    try {
      setIsPending(true);

      await clientApiFetch<{ message: string }>("/auth/logout", {
        method: "POST",
      });

      router.push("/");
      router.refresh();
    } finally {
      setIsPending(false);
    }
  }

  return (
    <Button
      className={logoutButtonClass}
      variant="default"
      onClick={handleLogout}
      disabled={isPending}
    >
      Logout
    </Button>
  );
}

export default LogoutButton;
