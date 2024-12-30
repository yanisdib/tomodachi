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

--- Creating tables (temporary structure)

-- PLATFORMS
CREATE TABLE IF NOT EXISTS public.platforms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    server VARCHAR(100),
    nitro_level SMALLINT,
    video_resolution VARCHAR(25),
    network_quality SMALLINT,
    access_link VARCHAR(255)
);

-- Adds comments to the columns of PLATFORMS table
COMMENT ON COLUMN platforms.name IS 'Name of the hosting platform';
COMMENT ON COLUMN platforms.server IS 'Name of the host''s server';
COMMENT ON COLUMN platforms.nitro_level IS 'Nitro level of the host''s Discord server';
COMMENT ON COLUMN platforms.video_resolution IS 'Resolution chosen for the event''s video stream';
COMMENT ON COLUMN platforms.network_quality IS 'Overall network quality of the host';
COMMENT ON COLUMN platforms.access_link IS 'Access link to the host room depending on the platform';

-- USERS
CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    CONSTRAINT username_format CHECK (username ~ '^[A-Za-z0-9_]+$'),
    email VARCHAR(255) NOT NULL,
    CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    birthdate BIGINT NOT NULL,
    verified BOOLEAN NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT
);

-- Adds comments to the columns of USERS table
COMMENT ON COLUMN users.username IS 'Username';
COMMENT ON COLUMN users.email IS 'Email address';
COMMENT ON COLUMN users.birthdate IS 'Birthdate as a raw epoch value';
COMMENT ON COLUMN users.verified IS 'Either user is verified or not';
COMMENT ON COLUMN users.created_at IS 'Account''s creation time as raw epoch value';
COMMENT ON COLUMN users.updated_at IS 'Account''s update time as raw epoch value';

-- LOBBIES
CREATE TABLE IF NOT EXISTS public.lobbies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    host_id INTEGER REFERENCES users(id),
    status SMALLINT NOT NULL,
    platform_id INTEGER REFERENCES platforms(id),
    start_at BIGINT,
    end_at BIGINT,
    access_rule SMALLINT NOT NULL,
    users_limit SMALLINT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT,
    CONSTRAINT check_start_end_time CHECK (start_at < end_at)
);

-- Adds comments to the columns of LOBBIES table
COMMENT ON COLUMN lobbies.name IS 'Name of the lobby';
COMMENT ON COLUMN lobbies.description IS 'Description of the lobby';
COMMENT ON COLUMN lobbies.host_id IS 'User hosting the current event';
COMMENT ON COLUMN lobbies.status IS 'Current status of the event';
COMMENT ON COLUMN lobbies.platform_id IS 'Platform used to host the event';
COMMENT ON COLUMN lobbies.start_at IS 'Event start time as epoch raw value';
COMMENT ON COLUMN lobbies.end_at IS 'Event end time as epoch raw value';
COMMENT ON COLUMN lobbies.access_rule IS 'Access mode defined to join a lobby';
COMMENT ON COLUMN lobbies.users_limit IS 'Limit of users accepted in a lobby';
COMMENT ON COLUMN lobbies.created_at IS 'Event''s creation time as raw epoch value';
COMMENT ON COLUMN lobbies.updated_at IS 'Event''s update time as raw epoch value';