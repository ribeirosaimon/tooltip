CREATE TABLE "test" (
                        ID         SERIAL PRIMARY KEY,
                        username   VARCHAR(255) NOT NULL,
                        password   VARCHAR(255) NOT NULL,
                        email      VARCHAR(255) NOT NULL UNIQUE,
                        first_name VARCHAR(255),
                        last_name  VARCHAR(255),
                        role       VARCHAR(50),
                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
