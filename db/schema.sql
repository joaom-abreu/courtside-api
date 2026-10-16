CREATE TABLE teams (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    city TEXT NOT NULL,
    abbreviation VARCHAR(3) NOT NULL UNIQUE,
    conference TEXT NOT NULL CHECK (conference IN ('East', 'West')),
    division TEXT NOT NULL
);
