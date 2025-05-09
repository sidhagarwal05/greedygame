CREATE TABLE campaigns (
    id TEXT PRIMARY KEY,
    image TEXT NOT NULL,
    cta TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE'))
);

CREATE TABLE targeting_rules (
    id SERIAL PRIMARY KEY,
    campaign_id TEXT REFERENCES campaigns(id) ON DELETE CASCADE,
    include_countries TEXT[],
    exclude_countries TEXT[],
    include_os TEXT[],
    exclude_os TEXT[],
    include_apps TEXT[],
    exclude_apps TEXT[]
);
