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

---- Table creation script for the Tomodachi platform

--- This script creates the necessary tables for the Tomodachi platform
--- The tables include:
--- 1. USERS: Stores information about users
--- 2. DISCORD_SERVERS: Stores information about Discord servers
--- 3. HUBS: Stores information about Hubs

-- Creates the USERS table
-- This table stores information about users
-- Users are the main participants in the platform
CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role SMALLINT NOT NULL,
    avatar_url VARCHAR(255),
    bio TEXT,
    location VARCHAR(100),
    website VARCHAR(255),
    birthdate BIGINT NOT NULL,
    verified BOOLEAN NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT,
    deleted_at BIGINT,
    CONSTRAINT pk_user PRIMARY KEY (id),
    CONSTRAINT fk_user FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT password_format CHECK (password ~* '^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{16,}$'),
    CONSTRAINT verified_check CHECK (verified IN (0, 1)),
    CONSTRAINT created_at_check CHECK (created_at <= EXTRACT(EPOCH FROM NOW())::BIGINT),
    CONSTRAINT updated_at_check CHECK (updated_at IS NULL OR updated_at >= created_at),
    CONSTRAINT deleted_at_check CHECK (deleted_at IS NULL OR deleted_at >= created_at),
    CONSTRAINT email_unique UNIQUE (email),
    CONSTRAINT username_unique UNIQUE (username),
    CONSTRAINT password_length CHECK (LENGTH(password) >= 16),
    CONSTRAINT role_check CHECK (role IN (0, 1, 2, 3)),
    CONSTRAINT birthdate_check CHECK (birthdate <= EXTRACT(EPOCH FROM NOW())::BIGINT),
    CONSTRAINT avatar_url_format CHECK (avatar_url ~* '^(https?:\/\/.*\.(?:png|jpg|jpeg|gif))$'),
    CONSTRAINT website_format CHECK (website ~* '^(https?:\/\/.*\.(?:png|jpg|jpeg|gif))$'),
    CONSTRAINT bio_length CHECK (LENGTH(bio) <= 500),
    CONSTRAINT location_length CHECK (LENGTH(location) <= 100),
    CONSTRAINT website_length CHECK (LENGTH(website) <= 255),
    CONSTRAINT avatar_url_length CHECK (LENGTH(avatar_url) <= 255),
    CONSTRAINT email_length CHECK (LENGTH(email) <= 255),
    CONSTRAINT username_length CHECK (LENGTH(username) <= 50),
    CONSTRAINT bio_format CHECK (bio ~* '^[A-Za-z0-9_.,!? ]*$'),
    CONSTRAINT location_format CHECK (location ~* '^[A-Za-z0-9_.,!? ]*$'),
    CONSTRAINT username_format CHECK (username ~ '^[A-Za-z0-9_]+$'),
    CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
);

-- Adds comments to the columns of USERS table
COMMENT ON COLUMN users.username IS 'Username';
COMMENT ON COLUMN users.email IS 'Email address';
COMMENT ON COLUMN users.birthdate IS 'Birthdate as a raw epoch value';
COMMENT ON COLUMN users.verified IS 'Either user is verified or not';
COMMENT ON COLUMN users.created_at IS 'Account''s creation time as raw epoch value';
COMMENT ON COLUMN users.updated_at IS 'Account''s update time as raw epoch value';


-- Creates the DISCORD_SERVERS table
-- This table stores information about Discord servers
-- Discord servers are the main platform for hosting events

CREATE TABLE IF NOT EXISTS public.discord_servers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nitro_level SMALLINT,
    icon_url VARCHAR(255),
    video_quality_mode SMALLINT(2),
    average_latency SMALLINT,
    shared_link VARCHAR(255),
    member_count SMALLINT,
    region VARCHAR(50),
    host_id INTEGER REFERENCES users(id),
    created_at BIGINT NOT NULL,
    updated_at BIGINT,
    deleted_at BIGINT,
    CONSTRAINT pk_discord_server PRIMARY KEY (id),
    CONSTRAINT fk_host FOREIGN KEY (host_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT check_video_quality_mode CHECK (video_quality_mode IN (1, 2, 3)),
    CONSTRAINT check_nitro_level CHECK (nitro_level IN (0, 1, 2, 3)),
    CONSTRAINT check_average_latency CHECK (average_latency >= 0),
    CONSTRAINT check_member_count CHECK (member_count >= 0)
);

-- Adds comments to the columns of DISCORD_SERVERS table
COMMENT ON COLUMN  discord_servers.name IS 'Name of the discord_server';
COMMENT ON COLUMN  discord_servers.nitro_level IS 'Nitro level of the discord_server';
COMMENT ON COLUMN  discord_servers.icon_url IS 'Icon URL of the discord_server';
COMMENT ON COLUMN  discord_servers.video_quality_mode IS 'Video quality mode of the discord_server';
COMMENT ON COLUMN  discord_servers.average_latency IS 'Average latency of the discord_server';
COMMENT ON COLUMN  discord_servers.shared_link IS 'Shared link of the discord_server';
COMMENT ON COLUMN  discord_servers.member_count IS 'Member count of the discord_server';
COMMENT ON COLUMN  discord_servers.region IS 'Region of the discord_server';
COMMENT ON COLUMN  discord_servers.host_id IS 'User hosting the current discord_server';
COMMENT ON COLUMN  discord_servers.created_at IS 'Platform''s creation time as raw epoch value';
COMMENT ON COLUMN  discord_servers.updated_at IS 'Platform''s update time as raw epoch value';
COMMENT ON COLUMN  discord_servers.deleted_at IS 'Platform''s deletion time as raw epoch value';


-- Creates the HUBS table
-- This table stores information about Hubs
-- Hubs are temporary events hosted by users on Discord servers

CREATE TABLE IF NOT EXISTS public.hubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    host_id INTEGER REFERENCES users(id),
    status SMALLINT NOT NULL,
    platform_id INTEGER REFERENCES discord_servers(id),
    start_at BIGINT,
    end_at BIGINT,
    access_rule SMALLINT NOT NULL,
    users_limit SMALLINT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT,
    deleted_at BIGINT,
    CONSTRAINT PK_hub PRIMARY KEY (id),
    CONSTRAINT fk_host FOREIGN KEY (host_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_platform FOREIGN KEY (platform_id) REFERENCES discord_servers(id) ON DELETE CASCADE,
    CONSTRAINT check_start_end_time CHECK (started_at < ended_at)
    CONSTRAINT check_status CHECK (status IN (0, 1, 2)),
    CONSTRAINT check_access_rule CHECK (access_rule IN (0, 1, 2)),
    CONSTRAINT check_users_limit CHECK (users_limit >= 0),
    CONSTRAINT check_name_length CHECK (LENGTH(name) <= 100),
    CONSTRAINT check_description_length CHECK (LENGTH(description) <= 500),
    CONSTRAINT check_name_format CHECK (name ~* '^[A-Za-z0-9_.,!? ]*$'),
    CONSTRAINT check_description_format CHECK (description ~* '^[A-Za-z0-9_.,!? ]*$'),
    CONSTRAINT check_host_id CHECK (host_id IS NOT NULL),
    CONSTRAINT check_platform_id CHECK (platform_id IS NOT NULL),
);

-- Adds comments to the columns of hubs table
COMMENT ON COLUMN hubs.name IS 'Name of the Hub';
COMMENT ON COLUMN hubs.description IS 'Description of the Hub';
COMMENT ON COLUMN hubs.host_id IS 'User hosting the current event';
COMMENT ON COLUMN hubs.status IS 'Current status of the event';
COMMENT ON COLUMN hubs.platform_id IS 'Platform used to host the event';
COMMENT ON COLUMN hubs.started_at IS 'Event start time as epoch raw value';
COMMENT ON COLUMN hubs.ended_at IS 'Event end time as epoch raw value';
COMMENT ON COLUMN hubs.access_rule IS 'Access mode defined to join a Hub';
COMMENT ON COLUMN hubs.users_limit IS 'Limit of users accepted in a Hub';
COMMENT ON COLUMN hubs.created_at IS 'Event''s creation time as raw epoch value';
COMMENT ON COLUMN hubs.updated_at IS 'Event''s update time as raw epoch value';
COMMENT ON COLUMN hubs.deleted_at IS 'Event''s deletion time as raw epoch value';

-- Creates the HUB_USERS table
-- This table stores information about users in Hubs
-- Users can join multiple Hubs

CREATE TABLE IF NOT EXISTS public.hub_users (
    hub_id INTEGER REFERENCES hubs(id) PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) PRIMARY KEY,
    joined_at BIGINT NOT NULL,
    left_at BIGINT,
    CONSTRAINT pk_hub_user PRIMARY KEY (hub_id, user_id),
    CONSTRAINT fk_hub FOREIGN KEY (hub_id) REFERENCES hubs(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    CONSTRAINT check_joined_at CHECK (joined_at <= EXTRACT(EPOCH FROM NOW())::BIGINT),
    CONSTRAINT check_left_at CHECK (left_at IS NULL OR left_at >= joined_at),
    CONSTRAINT check_hub_id CHECK (hub_id IS NOT NULL),
    CONSTRAINT check_user_id CHECK (user_id IS NOT NULL),
);

-- Adds comments to the columns of hub_users table
COMMENT ON COLUMN hub_users.hub_id IS 'Hub ID';
COMMENT ON COLUMN hub_users.user_id IS 'User ID';
COMMENT ON COLUMN hub_users.joined_at IS 'User joining time as raw epoch value';
COMMENT ON COLUMN hub_users.left_at IS 'User leaving time as raw epoch value';