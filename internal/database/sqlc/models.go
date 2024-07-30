// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID          int32       `json:"id"`
	Email       string      `json:"email"`
	Username    string      `json:"username"`
	PhoneNumber pgtype.Text `json:"phone_number"`
	// hashed password
	Password string `json:"password"`
	// role-based access: admin, seller, or buyer
	Role             string             `json:"role"`
	IsVerified       pgtype.Bool        `json:"is_verified"`
	CreatedAt        pgtype.Timestamptz `json:"created_at"`
	ResetToken       pgtype.Text        `json:"reset_token"`
	ResetTokenExpiry pgtype.Timestamptz `json:"reset_token_expiry"`
}
