CREATE TABLE IF NOT EXISTS submissions (
        id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        client_name VARCHAR(100) NOT NULL,
        client_email VARCHAR(100) NOT NULL,
        service_package VARCHAR(20) NOT NULL,
        project_goal TEXT NOT NULL,
        desired_timeline VARCHAR(50) NOT NULL,
        assets_provided VARCHAR(100) NOT NULL,

        readiness_status VARCHAR(20) NOT NULL
        CHECK(readiness_status IN('ready','missing_info')),
        missing_items TEXT[] NOT NULL DEFAULT '{}',
        ai_summary TEXT NOT NULL,
        recommended_next_action TEXT NOT NULL,

        admin_status VARCHAR(20) NOT NULL DEFAULT 'pending'
        CHECK(admin_status IN('pending','approved','rejected')),
        
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
        
)