CREATE TABLE IF NOT EXISTS doctors (
    id VARCHAR(50) PRIMARY KEY,
    last_name VARCHAR(100) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100) NOT NULL,
    positions TEXT[] NOT NULL,
    experience INTEGER NOT NULL,
    departments TEXT[] NOT NULL,
    is_active BOOLEAN DEFAULT true,
    available_for_appointment BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);