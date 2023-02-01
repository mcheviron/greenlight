CREATE TABLE IF NOT EXISTS
  movies (
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp(0)
    WITH
      TIME ZONE NOT NULL DEFAULT now(),
      title text NOT NULL,
      YEAR integer NOT NULL,
      runtime integer NOT NULL,
      genres text[] NOT NULL,
      VERSION integer NOT NULL DEFAULT 1
  );
