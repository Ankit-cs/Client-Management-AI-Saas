<h1 align="center">Client Onboarding SaaS</h1>

<p align="center">
  <strong>Production-ready Client Onboarding Portal & Workflow Automation</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Fiber-v3-00B3FF?style=for-the-badge&logo=gofiber&logoColor=white" alt="Fiber">
  <img src="https://img.shields.io/badge/Next.js-15+-000000?style=for-the-badge&logo=nextdotjs&logoColor=white" alt="Next.js">
  <img src="https://img.shields.io/badge/PostgreSQL-16+-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/n8n-Workflow-FF6C37?style=for-the-badge&logo=n8n&logoColor=white" alt="n8n">
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazonwebservices&logoColor=white" alt="AWS">
</p>

<hr />

## 💡 The Problem It Solves

Managing client intake in service industries is traditionally plagued by friction, disorganization, and delay:
*   **Information Intake Chaos**: Clients submit project details through fragmented channels (emails, chats, PDFs), leading to messy records.
*   **Manual Readiness Checks**: Teams waste valuable hours manually verifying if the client provided correct package details, target goals, timelines, and onboarding assets.
*   **Opaque Review Statuses**: Clients are left in the dark wondering if their onboarding request has been reviewed, approved, or rejected.

**Client Onboarding SaaS** solves this by establishing a centralized portal. It automates information verification using **n8n workflows** to evaluate client submissions for completeness (readiness check), registers missing requirements, and presents admins with a consolidated dashboard to review and approve projects in a single click.

---

## 📋 Prerequisites & System Requirements

Before running the application, make sure you have the following installed on your system:

| Requirement | Recommended Version | Purpose |
| :--- | :--- | :--- |
| **Go SDK** | `v1.26+` | Backend runtime & compilation |
| **Node.js** | `v18.x` or `v20.x` | Frontend React/Next.js execution |
| **Docker & Compose** | `Docker Desktop v24+` | Containerization of PostgreSQL and n8n |
| **Google Cloud Console** | Developer Account | Required to set up Google OAuth credentials |
| **Air** (Optional) | Latest | CLI hot-reloading tool for Go developers |

---

## 🛠️ Tech Stack & Key Dependencies

### 💰 Service Costs (Free vs Paid)
This project is built almost entirely on free and open-source tools, with the exception of the AI provider:

*   **Next.js (Frontend)**: Free & Open Source
*   **Go & Fiber (Backend)**: Free & Open Source
*   **PostgreSQL (Database)**: Free & Open Source
*   **Docker & Docker Compose**: Free & Open Source
*   **n8n (Workflow Engine)**: Free (Self-Hosted Community Edition)
*   **Google OAuth (Authentication)**: Free (Up to standard Google Cloud API limits)
*   **OpenAI API (AI Analysis)**: **Paid** (Pay-as-you-go based on token usage). *Note: You can swap this for a free local AI model like Ollama inside n8n if desired!*

### 🖥️ Frontend (Next.js Application)
A modern client-facing portal using TypeScript and CSS frameworks.
*   **Core Framework**: Next.js 16.2 (App Router)
*   **View Library**: React 19
*   **Styling & UI**: Tailwind CSS (v4) + Shadcn/UI (`@base-ui/react`, `class-variance-authority`, `tailwind-merge`)
*   **Icons**: `lucide-react`
*   **Animations**: `tw-animate-css`

### 🚀 Backend (Go Microservice)
A blazing fast, modular backend architecture handling REST requests, auth, database, and integrations.
*   **HTTP Engine**: `github.com/gofiber/fiber/v3`
*   **Database Pool**: `github.com/jackc/pgx/v5`
*   **Authentication**: `github.com/golang-jwt/jwt/v5` & `golang.org/x/oauth2`
*   **Dotenv Loader**: `github.com/joho/godotenv`

### 💾 Containerization & Database
*   **Database Engine**: PostgreSQL 16 (dockerized)
*   **Workflow Engine**: n8n (dockerized integration hub for executing AI readiness checks)

---

## 🗺️ System Flow & Architecture

```mermaid
sequenceDiagram
    participant C as Client (Next.js)
    participant A as Admin (Next.js)
    participant B as Backend (Go Fiber)
    participant N as Automation (n8n Workflow)
    participant DB as Database (PostgreSQL)

    C->>B: GET /auth/google (Login)
    B-->>C: Redirect & Set Auth Cookie (JWT)
    C->>B: POST /submissions (Intake Form)
    B->>N: Trigger Webhook (CheckReadiness)
    Note over N: n8n processes project details<br/>& runs AI assessment
    N-->>B: Return Status (ready/missing_info) + AI Summary
    B->>DB: Save User & Submission Details
    B-->>C: Success response
    A->>B: GET /admin/submissions (List all)
    B->>DB: Fetch submissions
    DB-->>B: Submissions list
    B-->>A: JSON Response
    A->>B: PATCH /admin/submissions/:id/status (Approve/Reject)
    B->>DB: Update admin_status
```

---

## 📂 Project Structure

```
c:\clientOnboarding\
├── docker-compose.yml       # PostgreSQL & n8n containers
├── README.md                # Project documentation
├── backend/                 # Go Fiber backend service
│   ├── cmd/
│   │   ├── api/             # Main server entrypoint (main.go)
│   │   └── migrate/         # SQL migration runner
│   ├── internal/
│   │   ├── config/          # Config parser using Env variables
│   │   ├── database/        # Migrator logic & schema migrations (SQL)
│   │   ├── https/           # Routes, Handlers, & Middlewares
│   │   ├── repositories/    # Database queries (User & Submission)
│   │   └── services/        # Business logic & external clients (n8n, Auth)
│   ├── go.mod
│   └── .air.toml            # Air auto-reload configurations
└── frontend/                # Next.js frontend application
    ├── src/
    │   ├── app/             # Page routes (admin, submissions, etc.)
    │   ├── components/      # UI components & shared site assets
    │   ├── lib/             # API request wrappers and configurations
    │   └── types/           # TS Interfaces
    └── package.json
```

---

## 💻 Important Commands Reference

Here are the crucial command lines needed to build, manage, and run the project:

### Services & Infrastructure (Docker)
Run these in the **project root directory** (`/`):
```bash
# Start Postgres & n8n in detached mode
docker-compose up -d

# Check running containers
docker-compose ps

# Stop all docker services
docker-compose down

# View logs of running services
docker-compose logs -f
```

### Backend Commands (Go)
Run these inside the `/backend` directory:
```bash
# Initialize and sync Go module dependencies
go mod tidy

# Run database migration scripts (creates tables)
go run cmd/migrate/main.go

# Run the live API server directly
go run cmd/api/main.go

# Start the Go server with hot-reload support (requires Air)
air
```

### Frontend Commands (Next.js)
Run these inside the `/frontend` directory:
```bash
# Install frontend package dependencies
npm install

# Run the client developer server (active on http://localhost:3000)
npm run dev

# Build the production bundle
npm run build

# Start the production server after building
npm run start
```

---

## 🔧 Detailed Configuration (.env)

### Backend Configuration (`/backend/.env`)
```env
APP_ENV=development
PORT=8080

FRONTEND_URL=http://localhost:3000
BACKEND_URL=http://localhost:8080

DATABASE_URL=postgres://postgres:Arbitrum123@127.0.0.1:5433/n8n_saas?sslmode=disable
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRES_IN_HOURS=168
AUTH_COOKIE_NAME=n8n_saas_token

# Google OAuth2 Credentials
GOOGLE_CLIENT_ID=your_google_client_id.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=your_google_client_secret
GOOGLE_REDIRECT_URL=http://localhost:8080/auth/google/callback

# N8N Integration
N8N_WEBHOOK_URL=http://localhost:5678/webhook-test/readiness-check
```

### Frontend Configuration (`/frontend/.env`)
```env
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
```

---

## 💾 Database Schema Reference

### `users`
*   `id`: UUID (Primary Key, defaults to random UUID)
*   `name`: VARCHAR (Client or Admin name)
*   `email`: VARCHAR (Unique, primary contact email)
*   `password`: VARCHAR (Encrypted password for local sign in)
*   `avatar_url`: VARCHAR (OAuth avatar image URL)
*   `role`: VARCHAR (`user` / `admin`, defaults to `user`)
*   `created_at`: TIMESTAMP

### `submissions`
*   `id`: UUID (Primary Key, defaults to random UUID)
*   `user_id`: UUID (Foreign Key to `users.id` with cascade deletion)
*   `client_name`: VARCHAR (Input Client Name)
*   `client_email`: VARCHAR (Input Client Email)
*   `service_package`: VARCHAR (Assigned package tier)
*   `project_goal`: TEXT (Detailed target goals)
*   `desired_timeline`: VARCHAR (Delivery timeline targets)
*   `assets_provided`: VARCHAR (Link/text indicating client assets)
*   `readiness_status`: VARCHAR (`ready` / `missing_info` - generated by AI review)
*   `missing_items`: TEXT[] (Array list of missing information items)
*   `ai_summary`: TEXT (Brief analysis of project goals generated by AI)
*   `recommended_next_action`: TEXT (AI suggestions for next steps)
*   `admin_status`: VARCHAR (`pending` / `approved` / `rejected`, defaults to `pending`)
*   `created_at`: TIMESTAMP
*   `updated_at`: TIMESTAMP
