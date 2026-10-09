CREATE TABLE teams (
                       id           BIGINT PRIMARY KEY,
                       name         TEXT        NOT NULL,
                       city         TEXT        NOT NULL,
                       abbreviation VARCHAR(3)  NOT NULL UNIQUE,
                       conference   VARCHAR(4)  NOT NULL CHECK (conference IN ('East', 'West')),
                       division     TEXT        NOT NULL,
                       created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE players (
                         id         BIGINT PRIMARY KEY,
                         first_name TEXT        NOT NULL,
                         last_name  TEXT        NOT NULL,
                         position   VARCHAR(5),
                         height_cm  SMALLINT    CHECK (height_cm > 0),
                         team_id    BIGINT      REFERENCES teams (id) ON DELETE SET NULL,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                         updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_players_team_id ON players (team_id);

CREATE TABLE seasons (
                         year SMALLINT PRIMARY KEY CHECK (year BETWEEN 1946 AND 2100)
    );

CREATE TABLE games (
                       id              BIGINT PRIMARY KEY,
                       season          SMALLINT    NOT NULL REFERENCES seasons (year),
                       game_date       DATE        NOT NULL,
                       home_team_id    BIGINT      NOT NULL REFERENCES teams (id),
                       visitor_team_id BIGINT      NOT NULL REFERENCES teams (id),
                       home_score      SMALLINT,
                       visitor_score   SMALLINT,
                       status          TEXT        NOT NULL DEFAULT 'scheduled'
                           CHECK (status IN ('scheduled', 'in_progress', 'final')),
                       postseason      BOOLEAN     NOT NULL DEFAULT false,
                       created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
                       CHECK (home_team_id <> visitor_team_id)
);

CREATE INDEX idx_games_season_date ON games (season, game_date);
CREATE INDEX idx_games_home_team_id ON games (home_team_id);
CREATE INDEX idx_games_visitor_team_id ON games (visitor_team_id);