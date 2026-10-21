package repository

import (
	"database/sql"
	"errors"

	"github.com/kittiphong064008-cyber/Go-Pokemon/pkg/pokemons/domain"
)

var (
	ErrNotFound      = errors.New("pokemon not found")
	ErrAlreadyExists = errors.New("pokemon with this ID or name already exists")
)

type Repository interface {
	Create(query domain.CreatePokemonQuery) (*domain.PokemonDetail, error)
	GetByID(id int) (*domain.PokemonDetail, error)
	GetAll() ([]domain.PokemonDetail, error)
	Update(id int, query domain.CreatePokemonQuery) (int, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(query domain.CreatePokemonQuery) (*domain.PokemonDetail, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	insertPokemonSQL := `
		INSERT INTO pokemon (id, name, height, weight, base_exp, hp, attack, defense, sp_atk, sp_def, speed, image_url, generation_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err = tx.Exec(insertPokemonSQL,
		query.ID, query.Name, query.Height, query.Weight, query.BaseExp,
		query.Hp, query.Attack, query.Defense, query.SpAtk, query.SpDef, query.Speed,
		query.ImageURL, query.GenerationID,
	)
	if err != nil {
		return nil, err
	}

	insertTypeSQL := `
		INSERT INTO pokemon_types (pokemon_id, type_id, slot)
		VALUES ($1, $2, $3)
	`
	for i, typeID := range query.TypeIDs {
		slot := i + 1
		if _, err := tx.Exec(insertTypeSQL, query.ID, typeID, slot); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return r.GetByID(query.ID)
}

func (r *repository) GetByID(id int) (*domain.PokemonDetail, error) {
	sqlQuery := `
		SELECT 
			p.id, p.name, p.height, p.weight, COALESCE(p.base_exp, 0),
			p.hp, p.attack, p.defense, p.sp_atk, p.sp_def, p.speed,
			COALESCE(p.image_url, ''), p.generation_id, p.created_at, p.updated_at,
			g.id, g.name, g.region
		FROM pokemon p
		LEFT JOIN generations g ON p.generation_id = g.id
		WHERE p.id = $1
	`

	var detail domain.PokemonDetail
	var genID sql.NullInt64
	var genName sql.NullString
	var genRegion sql.NullString

	err := r.db.QueryRow(sqlQuery, id).Scan(
		&detail.Pokemon.ID, &detail.Pokemon.Name, &detail.Pokemon.Height, &detail.Pokemon.Weight,
		&detail.Pokemon.BaseExp, &detail.Pokemon.Hp, &detail.Pokemon.Attack, &detail.Pokemon.Defense,
		&detail.Pokemon.SpAtk, &detail.Pokemon.SpDef, &detail.Pokemon.Speed,
		&detail.Pokemon.ImageURL, &detail.Pokemon.GenerationID,
		&detail.Pokemon.CreatedAt, &detail.Pokemon.UpdatedAt,
		&genID, &genName, &genRegion,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if genID.Valid {
		detail.Generation = &domain.Generation{
			ID:     int(genID.Int64),
			Name:   genName.String,
			Region: genRegion.String,
		}
	}

	typesQuery := `
		SELECT pt.pokemon_id, pt.type_id, pt.slot, t.name
		FROM pokemon_types pt
		JOIN types t ON pt.type_id = t.id
		WHERE pt.pokemon_id = $1
		ORDER BY pt.slot ASC
	`
	rows, err := r.db.Query(typesQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pt domain.PokemonType
		if err := rows.Scan(&pt.PokemonID, &pt.TypeID, &pt.Slot, &pt.TypeName); err != nil {
			return nil, err
		}
		detail.Types = append(detail.Types, pt)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &detail, nil
}

func (r *repository) GetAll() ([]domain.PokemonDetail, error) {
	sqlQuery := `
		SELECT 
			p.id, p.name, p.height, p.weight, COALESCE(p.base_exp, 0),
			p.hp, p.attack, p.defense, p.sp_atk, p.sp_def, p.speed,
			COALESCE(p.image_url, ''), p.generation_id, p.created_at, p.updated_at,
			g.id, g.name, g.region
		FROM pokemon p
		LEFT JOIN generations g ON p.generation_id = g.id
		ORDER BY p.id ASC
	`

	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.PokemonDetail
	for rows.Next() {
		var detail domain.PokemonDetail
		var genID sql.NullInt64
		var genName sql.NullString
		var genRegion sql.NullString

		err := rows.Scan(
			&detail.Pokemon.ID, &detail.Pokemon.Name, &detail.Pokemon.Height, &detail.Pokemon.Weight,
			&detail.Pokemon.BaseExp, &detail.Pokemon.Hp, &detail.Pokemon.Attack, &detail.Pokemon.Defense,
			&detail.Pokemon.SpAtk, &detail.Pokemon.SpDef, &detail.Pokemon.Speed,
			&detail.Pokemon.ImageURL, &detail.Pokemon.GenerationID,
			&detail.Pokemon.CreatedAt, &detail.Pokemon.UpdatedAt,
			&genID, &genName, &genRegion,
		)
		if err != nil {
			return nil, err
		}

		if genID.Valid {
			detail.Generation = &domain.Generation{
				ID:     int(genID.Int64),
				Name:   genName.String,
				Region: genRegion.String,
			}
		}

		results = append(results, detail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	for i := range results {
		typesQuery := `
			SELECT pt.pokemon_id, pt.type_id, pt.slot, t.name
			FROM pokemon_types pt
			JOIN types t ON pt.type_id = t.id
			WHERE pt.pokemon_id = $1
			ORDER BY pt.slot ASC
		`
		tRows, err := r.db.Query(typesQuery, results[i].Pokemon.ID)
		if err != nil {
			return nil, err
		}
		for tRows.Next() {
			var pt domain.PokemonType
			if err := tRows.Scan(&pt.PokemonID, &pt.TypeID, &pt.Slot, &pt.TypeName); err != nil {
				tRows.Close()
				return nil, err
			}
			results[i].Types = append(results[i].Types, pt)
		}
		if err := tRows.Err(); err != nil {
			tRows.Close()
			return nil, err
		}
		tRows.Close()
	}

	return results, nil
}

func (r *repository) Update(id int, query domain.CreatePokemonQuery) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	sqlQuery := `
		UPDATE pokemon 
		SET name = $1, height = $2, weight = $3, base_exp = $4,
		    hp = $5, attack = $6, defense = $7, sp_atk = $8, sp_def = $9, speed = $10,
		    image_url = $11, generation_id = $12, updated_at = now()
		WHERE id = $13
	`
	result, err := tx.Exec(sqlQuery,
		query.Name, query.Height, query.Weight, query.BaseExp,
		query.Hp, query.Attack, query.Defense, query.SpAtk, query.SpDef, query.Speed,
		query.ImageURL, query.GenerationID, id,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAffected == 0 {
		return 0, ErrNotFound
	}

	if len(query.TypeIDs) > 0 {
		if _, err := tx.Exec(`DELETE FROM pokemon_types WHERE pokemon_id = $1`, id); err != nil {
			return 0, err
		}
		insertTypeSQL := `INSERT INTO pokemon_types (pokemon_id, type_id, slot) VALUES ($1, $2, $3)`
		for i, typeID := range query.TypeIDs {
			slot := i + 1
			if _, err := tx.Exec(insertTypeSQL, id, typeID, slot); err != nil {
				return 0, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}

func (r *repository) Delete(id int) error {
	sqlQuery := `DELETE FROM pokemon WHERE id = $1`
	result, err := r.db.Exec(sqlQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
