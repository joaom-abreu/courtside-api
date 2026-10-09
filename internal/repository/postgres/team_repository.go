package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/joaom-abreu/courtside-api/internal/domain"
	"github.com/joaom-abreu/courtside-api/internal/repository/postgres/sqlc"
)

type TeamRepository struct {
	queries *sqlc.Queries
}

func NewTeamRepository(db sqlc.DBTX) *TeamRepository {
	return &TeamRepository{queries: sqlc.New(db)}
}

func (r *TeamRepository) List(ctx context.Context, conference domain.Conference) ([]domain.Team, error) {
	var filter *string
	if conference != "" {
		value := string(conference)
		filter = &value
	}

	rows, err := r.queries.ListTeams(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("listing teams: %w", err)
	}

	teams := make([]domain.Team, 0, len(rows))
	for _, row := range rows {
		teams = append(teams, toDomainTeam(row.ID, row.Name, row.City, row.Abbreviation, row.Conference, row.Division))
	}

	return teams, nil
}

func (r *TeamRepository) GetByID(ctx context.Context, id int64) (domain.Team, error) {
	row, err := r.queries.GetTeamByID(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.Team{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.Team{}, fmt.Errorf("getting team %d: %w", id, err)
	}

	return toDomainTeam(row.ID, row.Name, row.City, row.Abbreviation, row.Conference, row.Division), nil
}

func (r *TeamRepository) Upsert(ctx context.Context, team domain.Team) error {
	err := r.queries.UpsertTeam(ctx, sqlc.UpsertTeamParams{
		ID:           team.ID,
		Name:         team.Name,
		City:         team.City,
		Abbreviation: team.Abbreviation,
		Conference:   string(team.Conference),
		Division:     team.Division,
	})
	if err != nil {
		return fmt.Errorf("upserting team %d: %w", team.ID, err)
	}

	return nil
}

func toDomainTeam(id int64, name, city, abbreviation, conference, division string) domain.Team {
	return domain.Team{
		ID:           id,
		Name:         name,
		City:         city,
		Abbreviation: abbreviation,
		Conference:   domain.Conference(conference),
		Division:     division,
	}
}
