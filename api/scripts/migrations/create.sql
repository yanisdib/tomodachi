-- Creates new database and assigns the owner
CREATE DATABASE tomodachi;

-- Connects to the newly created database
\c tomodachi 

-- Creates the public schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS public AUTHORIZATION admin;

-- Sets the search path for the database
ALTER DATABASE tomodachi SET search_path = public;

-- Grants all privileges to the admin user
GRANT ALL PRIVILEGES ON DATABASE tomodachi TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin;

-- USERS Table
CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE CHECK (username ~ '^[A-Za-z0-9_]+$'),
    email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    password VARCHAR(74) NOT NULL CHECK (LENGTH(password) >= 16),
    role VARCHAR(20) NOT NULL CHECK (role IN ('default', 'admin')),
    avatar_url VARCHAR(255),
    bio TEXT,
    location VARCHAR(100),
    website VARCHAR(255),
    birthdate BIGINT NOT NULL CHECK (birthdate <= EXTRACT(EPOCH FROM NOW())::BIGINT),
    verified BOOLEAN NOT NULL,
    created_at BIGINT NOT NULL CHECK (created_at <= EXTRACT(EPOCH FROM NOW())::BIGINT),
    updated_at BIGINT CHECK (updated_at IS NULL OR updated_at >= created_at)
);

-- USERS Column Comments
COMMENT ON COLUMN users.username IS 'Username';
COMMENT ON COLUMN users.email IS 'Email address';
COMMENT ON COLUMN users.birthdate IS 'Birthdate as a raw epoch value';
COMMENT ON COLUMN users.verified IS 'Either user is verified or not';
COMMENT ON COLUMN users.created_at IS 'Account''s creation time as raw epoch value';
COMMENT ON COLUMN users.updated_at IS 'Account''s update time as raw epoch value';

-- Insert sample admin user
INSERT INTO public.users (
    username, email, password, role, avatar_url, bio, location, website, 
    birthdate, verified, created_at
) VALUES (
    'daimao',
    'daimao@tomodachi.com',
    '$2a$12$mwGpOjq4cdU2hUwOiXRg6eD9ZGhItFLoKku7ob4LbQJAhU52IXo3y',
    'admin',
    'https://example.com/avatars/daimao.png',
    'Loves organizing DBZ tournaments and playing competitive fighters!',
    'West City',
    'https://twitch.tv/daimao',
    1747872000,
    true,
    EXTRACT(EPOCH FROM NOW())::BIGINT
);

-- DISCORD_SERVERS Table
CREATE TABLE IF NOT EXISTS public.discord_servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nitro_level SMALLINT CHECK (nitro_level IN (0, 1, 2, 3)),
    icon_url VARCHAR(255),
    video_quality_mode SMALLINT CHECK (video_quality_mode IN (1, 2, 3)),
    average_latency SMALLINT CHECK (average_latency >= 0),
    shared_link VARCHAR(255),
    member_count SMALLINT CHECK (member_count >= 0),
    region VARCHAR(50),
    host_id INTEGER REFERENCES users(id),
    created_at BIGINT NOT NULL,
    updated_at BIGINT,
    closed_at BIGINT
);

-- DISCORD_SERVERS Column Comments
COMMENT ON COLUMN discord_servers.name IS 'Name of the discord_server';
COMMENT ON COLUMN discord_servers.nitro_level IS 'Nitro level of the discord_server';
COMMENT ON COLUMN discord_servers.icon_url IS 'Icon URL of the discord_server';
COMMENT ON COLUMN discord_servers.video_quality_mode IS 'Video quality mode of the discord_server';
COMMENT ON COLUMN discord_servers.average_latency IS 'Average latency of the discord_server';
COMMENT ON COLUMN discord_servers.shared_link IS 'Shared link of the discord_server';
COMMENT ON COLUMN discord_servers.member_count IS 'Member count of the discord_server';
COMMENT ON COLUMN discord_servers.region IS 'Region of the discord_server';
COMMENT ON COLUMN discord_servers.host_id IS 'User hosting the current discord_server';
COMMENT ON COLUMN discord_servers.created_at IS 'Platform''s creation time as raw epoch value';
COMMENT ON COLUMN discord_servers.updated_at IS 'Platform''s update time as raw epoch value';
COMMENT ON COLUMN discord_servers.closed_at IS 'Platform''s deletion time as raw epoch value';

-- Insert sample server
INSERT INTO public.discord_servers (
    name, nitro_level, video_quality_mode, average_latency, shared_link,
    member_count, region, host_id, created_at
) VALUES (
    'Dragon Ball Sparking Zero - Worldwide Tournament of Power 25',
    2,
    2,
    45,
    'https://discord.gg/sparkingzeroww',
    150,
    'NA-East',
    1,
    EXTRACT(EPOCH FROM NOW())::BIGINT
);

CREATE TABLE IF NOT EXISTS public.hubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT CHECK (LENGTH(description) <= 500),
    icon_url VARCHAR(255),
    tags TEXT[],
    event_type VARCHAR(50),
    languages TEXT[],
    user_limit SMALLINT CHECK (user_limit >= 0),
    access_rule VARCHAR(20) NOT NULL CHECK (
        access_rule IN (
            'public', 'private', 'code', 'request', 'password', 'invite'
        )
    ),
    password VARCHAR(100),
    start_at BIGINT,
    end_at BIGINT,
    active BOOLEAN NOT NULL DEFAULT false,
    open BOOLEAN NOT NULL DEFAULT false,
    created_at BIGINT NOT NULL,
    updated_at BIGINT,
    closed_at BIGINT,
    server_id BIGINT NOT NULL REFERENCES discord_servers(id) ON DELETE CASCADE,
    host_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);


-- HUBS Column Comments
COMMENT ON COLUMN hubs.name IS 'Name of the Hub';
COMMENT ON COLUMN hubs.description IS 'Description of the Hub';
COMMENT ON COLUMN hubs.icon_url IS 'Optional icon URL for the Hub';
COMMENT ON COLUMN hubs.tags IS 'String array of tags categorizing the Hub';
COMMENT ON COLUMN hubs.event_type IS 'Type or theme of the Hub event';
COMMENT ON COLUMN hubs.languages IS 'Languages supported in this Hub';
COMMENT ON COLUMN hubs.user_limit IS 'Maximum number of users allowed in the Hub';
COMMENT ON COLUMN hubs.access_rule IS 'Access rule for joining the Hub';
COMMENT ON COLUMN hubs.password IS 'Optional access password for private or restricted Hubs';
COMMENT ON COLUMN hubs.start_at IS 'Epoch timestamp for Hub start time';
COMMENT ON COLUMN hubs.end_at IS 'Epoch timestamp for Hub end time';
COMMENT ON COLUMN hubs.is_active IS 'Whether the Hub is currently active';
COMMENT ON COLUMN hubs.is_open IS 'Whether the Hub is open for joining';
COMMENT ON COLUMN hubs.created_at IS 'Epoch timestamp when the Hub was created';
COMMENT ON COLUMN hubs.updated_at IS 'Epoch timestamp when the Hub was last updated';
COMMENT ON COLUMN hubs.closed_at IS 'Epoch timestamp when the Hub was closed';
COMMENT ON COLUMN hubs.server_id IS 'Foreign key to Discord server hosting this Hub';
COMMENT ON COLUMN hubs.host_id IS 'User ID of the host who created the Hub';


-- HUB_USERS Table
CREATE TABLE IF NOT EXISTS public.hub_users (
    hub_id INTEGER NOT NULL REFERENCES hubs(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at BIGINT NOT NULL CHECK (joined_at <= EXTRACT(EPOCH FROM NOW())::BIGINT),
    left_at BIGINT CHECK (left_at IS NULL OR left_at >= joined_at),
    PRIMARY KEY (hub_id, user_id)
);

-- HUB_USERS Column Comments
COMMENT ON COLUMN hub_users.hub_id IS 'Hub ID';
COMMENT ON COLUMN hub_users.user_id IS 'User ID';
COMMENT ON COLUMN hub_users.joined_at IS 'User joining time as raw epoch value';
COMMENT ON COLUMN hub_users.left_at IS 'User leaving time as raw epoch value';
