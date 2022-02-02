CREATE TABLE IF NOT EXISTS jobs (
  position text NOT NULL,
  organization text NOT NULL,
  url text,
  description text,
  email text NOT NULL,
  published_at timestamp DEFAULT current_timestamp
);
