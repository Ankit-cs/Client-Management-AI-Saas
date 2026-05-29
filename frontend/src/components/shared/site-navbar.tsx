import { cn } from "@/lib/utils";
import { buttonVariants } from "../ui/button";
import { User } from "@/types/auth";
import Link from "next/link";
import { publicApiBaseUrl } from "@/lib/config";
import LogoutButton from "./logout-button";

const brandLinkClass = "flex items-center gap-3 group";

const brandLogoClass =
  "flex h-10 w-10 items-center justify-center border border-white/20 bg-black text-sm font-bold text-white rounded-xl group-hover:scale-[1.05] transition-all duration-200";

const brandTextWrapperClass = "grid gap-0.5";

const brandTitleClass = "text-base font-bold tracking-tight text-white font-inter group-hover:text-zinc-300 transition-colors duration-200";

const headerClass = "sticky top-0 z-40 border-b border-white/10 bg-black/80 backdrop-blur-md";

const navbarContainerClass =
  "mx-auto flex max-w-7xl items-center justify-between gap-4 px-6 py-4";

const navActionsClass = "flex items-center gap-4";

const workspaceButtonClass = cn(
  "inline-flex items-center justify-center h-10 rounded-xl px-4 text-sm font-medium text-white hover:text-zinc-300 hover:bg-white/5 transition-all duration-200 font-inter",
);

const loginButtonClass = cn(
  "inline-flex items-center justify-center h-10 bg-white hover:bg-zinc-200 hover:scale-[1.03] px-5 text-sm font-bold text-black shadow-[0_4px_12px_rgba(255,255,255,0.15)] transition-all duration-200 rounded-xl font-inter",
);

function SiteNavbar({ user }: { user: User | null }) {
  const workspaceHref = user?.role === "admin" ? "/admin" : "/submissions";

  return (
    <header className={headerClass}>
      <div className={navbarContainerClass}>
        <Link className={brandLinkClass} href="/">
          <span className={brandLogoClass}>OA</span>
          <div className={brandTextWrapperClass}>
            <span className={brandTitleClass}>OnboardingAI</span>
          </div>
        </Link>

        {user ? (
          <div className={navActionsClass}>
            <Link href={workspaceHref} className={workspaceButtonClass}>
              Workspace
            </Link>
            <LogoutButton />
          </div>
        ) : (
          <Link
            href={`${publicApiBaseUrl}/auth/google`}
            className={loginButtonClass}
          >
            Login
          </Link>
        )}
      </div>
    </header>
  );
}

export default SiteNavbar;
