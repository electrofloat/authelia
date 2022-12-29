CREATE TABLE IF NOT EXISTS one_time_password (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    signature VARCHAR(128) NOT NULL,
    iat TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    issued_ip VARCHAR(39) NOT NULL,
    exp TIMESTAMP NOT NULL,
    username VARCHAR(100) NOT NULL,
    intent VARCHAR(100) NOT NULL,
    consumed TIMESTAMP NULL DEFAULT NULL,
    consumed_ip VARCHAR(39) NULL DEFAULT NULL,
    password BLOB NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

CREATE UNIQUE INDEX one_time_password_signature ON one_time_password (signature);
CREATE INDEX one_time_password_lookup ON one_time_password (signature, username);

CREATE TABLE IF NOT EXISTS user_elevated_session (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_ip VARCHAR(39) NOT NULL,
    method VARCHAR(10) NOT NULL,
    method_id INTEGER NULL,
    expires TIMESTAMP NOT NULL,
    username VARCHAR(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;

CREATE INDEX user_elevated_session_username ON user_elevated_session (username);