-- creation of projects table
CREATE TABLE IF NOT EXISTS projects (
    project_id VARCHAR(225) NOT NULL,
    name VARCHAR(225) NOT NULL,
    PRIMARY KEY (project_id)
);

-- creation of flags table
CREATE TABLE IF NOT EXISTS flags (
    flag_id VARCHAR(225) NOT NULL,
    name VARCHAR(225) NOT NULL,
    value BOOLEAN NOT NULL,
    canary_setting INT DEFAULT 100 NOT NULL CHECK (canary_setting BETWEEN 1 AND 100),
    project_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (flag_id)
    );

-- add relationship between projects table and flags table
ALTER TABLE flags
    ADD CONSTRAINT fk_project
        FOREIGN KEY (project_id)
            REFERENCES projects(project_id)
            ON DELETE CASCADE;
