-- name: ListTeams :many
-- Returns every team, optionally filtered by conference ('East' or 'West').
SELECT id, name, city, abbreviation, conference, division
FROM teams
WHERE sqlc.narg('conference')::text IS NULL
   OR conference = sqlc.narg('conference')::text
ORDER BY city, name;

-- name: GetTeamByID :one
SELECT id, name, city, abbreviation, conference, division
FROM teams
WHERE id = $1;

-- name: UpsertTeam :exec
-- Inserts a team or, if the id already exists, updates it.
INSERT INTO teams (id, name, city, abbreviation, conference, division)
VALUES ($1, $2, $3, $4, $5, $6)
    ON CONFLICT (id) DO UPDATE SET
    name         = EXCLUDED.name,
                            city         = EXCLUDED.city,
                            abbreviation = EXCLUDED.abbreviation,
                            conference   = EXCLUDED.conference,
                            division     = EXCLUDED.division,
                            updated_at   = now();