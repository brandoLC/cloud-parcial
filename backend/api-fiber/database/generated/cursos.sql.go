// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: cursos.sql

package generated

import (
	"context"
	"database/sql"
)

const createCurso = `-- name: CreateCurso :execresult
INSERT INTO
    cursos (codigo, nombre, horario, ciclo)
VALUES
    (?, ?, ?, ?)
`

type CreateCursoParams struct {
	Codigo  int32
	Nombre  string
	Horario string
	Ciclo   string
}

func (q *Queries) CreateCurso(ctx context.Context, arg CreateCursoParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCurso,
		arg.Codigo,
		arg.Nombre,
		arg.Horario,
		arg.Ciclo,
	)
}

const deleteCurso = `-- name: DeleteCurso :exec
DELETE FROM cursos
WHERE
    codigo = ?
`

func (q *Queries) DeleteCurso(ctx context.Context, codigo int32) error {
	_, err := q.db.ExecContext(ctx, deleteCurso, codigo)
	return err
}

const getCurso = `-- name: GetCurso :one
SELECT
    codigo, nombre, horario, ciclo, created_at
FROM
    cursos
WHERE
    codigo = ?
LIMIT
    1
`

func (q *Queries) GetCurso(ctx context.Context, codigo int32) (Curso, error) {
	row := q.db.QueryRowContext(ctx, getCurso, codigo)
	var i Curso
	err := row.Scan(
		&i.Codigo,
		&i.Nombre,
		&i.Horario,
		&i.Ciclo,
		&i.CreatedAt,
	)
	return i, err
}

const listCursos = `-- name: ListCursos :many
SELECT
    codigo, nombre, horario, ciclo, created_at
FROM
    cursos
ORDER BY
    nombre
`

func (q *Queries) ListCursos(ctx context.Context) ([]Curso, error) {
	rows, err := q.db.QueryContext(ctx, listCursos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Curso
	for rows.Next() {
		var i Curso
		if err := rows.Scan(
			&i.Codigo,
			&i.Nombre,
			&i.Horario,
			&i.Ciclo,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCurso = `-- name: UpdateCurso :exec
UPDATE cursos
SET
    nombre = ?,
    horario = ?,
    ciclo = ?
WHERE
    codigo = ?
`

type UpdateCursoParams struct {
	Nombre  string
	Horario string
	Ciclo   string
	Codigo  int32
}

func (q *Queries) UpdateCurso(ctx context.Context, arg UpdateCursoParams) error {
	_, err := q.db.ExecContext(ctx, updateCurso,
		arg.Nombre,
		arg.Horario,
		arg.Ciclo,
		arg.Codigo,
	)
	return err
}
