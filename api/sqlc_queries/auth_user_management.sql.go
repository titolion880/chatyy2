// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: auth_user_management.sql

package sqlc_queries

import (
	"context"
)

const getRateLimit = `-- name: GetRateLimit :one
SELECT rate_limit AS rate_limit
FROM auth_user_management
WHERE user_id = $1
`

// GetRateLimit retrieves the rate limit for a user from the auth_user_management table.
// If no rate limit is set for the user, it returns the default rate limit of 100.
func (q *Queries) GetRateLimit(ctx context.Context, userID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRateLimit, userID)
	var rate_limit int32
	err := row.Scan(&rate_limit)
	return rate_limit, err
}
