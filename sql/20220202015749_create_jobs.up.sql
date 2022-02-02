CREATE TABLE IF NOT EXISTS jobs (
  id SERIAL PRIMARY KEY,
  position TEXT NOT NULL,
  organization TEXT NOT NULL,
  url TEXT,
  description TEXT,
  email TEXT NOT NULL,
  published_at TIMESTAMP DEFAULT current_timestamp
);