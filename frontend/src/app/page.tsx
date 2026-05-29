import SiteNavbar from "@/components/shared/site-navbar";
import { buttonVariants } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { getUserInfo } from "@/lib/api/auth";
import { publicApiBaseUrl } from "@/lib/config";
import { cn } from "@/lib/utils";
import Link from "next/link";
import { CheckCircle2, AlertCircle } from "lucide-react";
import Footer from "@/components/shared/footer";

const workflowSteps = [
  {
    title: "Submit onboarding details",
    description: "Provide your project goal, target audience, technical stack, timeline, assets, and custom notes in a structured form.",
  },
  {
    title: "Get an AI readiness review",
    description: "Our AI immediately analyzes your submission to determine readiness, flags missing items, and suggests next actions.",
  },
  {
    title: "Wait for admin decision",
    description: "The agency review team approves or rejects the onboarding request. View the real-time status in your client workspace.",
  },
];

const benefits = [
  {
    title: "Clear Intake",
    description:
      "Collect the service, project goal, timeline, assets, and notes in one structured onboarding flow.",
  },
  {
    title: "Faster Review",
    description:
      "The team gets a readiness summary, missing items, next action, and suggested follow-up before making a decision.",
  },
  {
    title: "Simple Approval Flow",
    description:
      "Clients can see whether the agency approved or rejected their onboarding request without any extra checkout step.",
  },
];

export default async function Home() {
  const user = await getUserInfo();
  const workspaceHref = user?.role === "admin" ? "/admin" : "/submissions";
  const loginHref = `${publicApiBaseUrl}/auth/google`;
  const viewProcessHref = user ? workspaceHref : "/submissions";

  return (
    <>
      <SiteNavbar user={user} />

      <main className="min-h-screen w-full flex flex-col items-center justify-center px-4 sm:px-6 lg:px-8 relative bg-black text-white overflow-hidden pb-24">
        {/* Grid Background */}
        <div className="absolute left-0 right-0 top-0 h-[120vh] opacity-60 z-0 hero-grid-bg" aria-hidden="true" />

        {/* Hero Section */}
        <div className="max-w-5xl w-full mx-auto text-center mt-28 sm:mt-36 relative z-10">
          <div className="flex justify-center mb-6">
            <div className="inline-flex items-center gap-2 rounded-full border border-white/10 bg-white/5 px-4 py-1.5 text-xs sm:text-sm text-white/80 font-medium font-inter">
              <span className="w-2 h-2 rounded-full bg-white animate-pulse" />
              <span>AI-Powered Client Onboarding SaaS</span>
            </div>
          </div>

          <h1 className="text-4xl sm:text-6xl lg:text-7xl font-inter font-bold leading-[1.15] tracking-tight">
            Client onboarding
            <br />
            <span className="text-white">in record time</span>
          </h1>

          <p className="max-w-2xl mt-6 mx-auto text-base sm:text-lg font-inter font-medium text-gray-300 leading-relaxed">
            Verify readiness, flags missing information, and automate client collection with a simple, secure onboarding system powered by AI.
          </p>

          <div className="flex mt-10 flex-col sm:flex-row gap-4 justify-center items-center">
            {user ? (
              <Link 
                href={workspaceHref}
                className="inline-flex items-center justify-center px-8 py-3.5 text-base font-bold bg-white text-black rounded-xl hover:scale-[1.03] transition-all duration-200 shadow-[0_6px_20px_rgba(255,255,255,0.15)] font-inter"
              >
                Open Workspace
                <svg className="ml-2 w-4 h-4 self-center" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                </svg>
              </Link>
            ) : (
              <a 
                href={loginHref}
                className="inline-flex items-center justify-center px-8 py-3.5 text-base font-bold bg-white text-black rounded-xl hover:scale-[1.03] transition-all duration-200 shadow-[0_6px_20px_rgba(255,255,255,0.15)] font-inter"
              >
                Start Onboarding
                <svg className="ml-2 w-4 h-4 self-center" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                </svg>
              </a>
            )}
            
            <Link 
              href={viewProcessHref}
              className="group cursor-pointer relative inline-flex items-center justify-center px-8 py-3 text-base font-medium text-white bg-transparent border-2 border-[#414141] hover:scale-[1.03] hover:border-white transition-all duration-200 rounded-xl h-12.5 font-inter"
            >
              View Process
            </Link>
          </div>
        </div>

        {/* Benefits Section */}
        <section className="w-full max-w-6xl mx-auto mt-28 relative z-10 px-4">
          <div className="text-center mb-16">
            <h2 className="font-inter font-bold text-3xl sm:text-4xl text-white mb-4">
              Why choose OnboardAI?
            </h2>
            <p className="text-base sm:text-lg text-gray-400 max-w-2xl mx-auto">
              Built for modern agencies to eliminate back-and-forth emails and accelerate project kickoff.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {benefits.map((benefit, index) => (
              <div
                key={index}
                className="rounded-2xl border border-white/10 bg-black/60 p-8 backdrop-blur-sm shadow-[0_0_0_1px_rgba(255,255,255,0.04)_inset,0_10px_40px_rgba(0,0,0,0.35)] hover:border-white/20 transition-all duration-300 flex flex-col justify-between"
              >
                <div>
                  <div className="w-12 h-12 rounded-xl bg-white/10 flex items-center justify-center mb-6">
                    <CheckCircle2 className="w-6 h-6 text-white" />
                  </div>
                  <h3 className="font-dm-sans font-bold text-2xl text-white mb-3">
                    {benefit.title}
                  </h3>
                  <p className="text-base font-dm-sans text-gray-300 font-medium leading-relaxed">
                    {benefit.description}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </section>

        {/* How It Works Section */}
        <section className="w-full max-w-6xl mx-auto mt-28 relative z-10 px-4">
          <div className="grid grid-cols-1 lg:grid-cols-12 gap-10 lg:gap-16 items-start">
            <div className="lg:col-span-5">
              <div className="inline-flex items-center gap-2 rounded-full border border-white/10 bg-white/5 px-4 py-1 text-xs sm:text-sm text-white/80 font-medium font-inter mb-4">
                <span>The Workflow</span>
              </div>
              <h2 className="font-inter font-bold text-3xl sm:text-4xl text-white leading-tight">
                Simple steps to onboarding success
              </h2>
              <p className="mt-4 text-base sm:text-lg text-gray-400 leading-relaxed">
                Our onboarding flows ensure that you have every single asset, stack specification, and timeline agreement ready before developers write a line of code.
              </p>
            </div>

            <div className="lg:col-span-7 space-y-6">
              {workflowSteps.map((step, idx) => (
                <div
                  key={idx}
                  className="flex gap-4 sm:gap-6 rounded-2xl border border-white/10 bg-black/60 p-6 backdrop-blur-sm shadow-[0_0_0_1px_rgba(255,255,255,0.04)_inset] hover:border-white/20 transition-colors duration-300"
                >
                  <div className="flex items-center justify-center h-10 w-10 shrink-0 rounded-full bg-white text-black font-inter font-bold text-base shadow-[0_4px_12px_rgba(255,255,255,0.15)]">
                    {idx + 1}
                  </div>
                  <div>
                    <h3 className="font-dm-sans font-bold text-lg text-white">{step.title}</h3>
                    <p className="mt-1 text-sm font-dm-sans text-gray-400 leading-relaxed">{step.description}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </>
  );
}

