CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    points INT DEFAULT 0,
    wins INT DEFAULT 0,
    draws INT DEFAULT 0,
    losses INT DEFAULT 0,
    goals_for INT DEFAULT 0,
    goals_against INT DEFAULT 0
);

CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    home_team_id INT REFERENCES teams(id),
    away_team_id INT REFERENCES teams(id),
    home_goals INT,
    away_goals INT,
    week INT
);
