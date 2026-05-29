"use client";

import Link from "next/link";
import { FolderGit2 } from "lucide-react";

export default function Footer() {
  return (
    <footer className="text-white/80 font-inter border-t border-white/10 w-full bg-black">
      <div className="max-w-[86rem] mx-auto px-6 py-8">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-6">
          <div className="flex flex-col gap-2 text-left">
            <div className="flex gap-3 items-center">
              <div className="flex-shrink-0 text-white">
                <FolderGit2 className="w-8 h-8" />
              </div>
              <span className="font-semibold text-2xl font-inter text-white">
                OnboardingAI
              </span>
            </div>
            <p className="text-sm font-inter font-medium text-white/55 mt-1">
              © {new Date().getFullYear()} OnboardingAI. All rights reserved.
            </p>
          </div>
          
          <div className="flex items-center gap-4">
            <Link
              href="https://github.com/Ankit-cs/Client-Management-AI-Saas"
              target="_blank"
              rel="noreferrer"
              className="text-white/60 hover:text-white transition-colors"
            >
              <svg className="w-7 h-7" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4" />
                <path d="M9 18c-4.51 2-5-2-7-2" />
              </svg>
            </Link>
          </div>
        </div>
      </div>
    </footer>
  );
}
