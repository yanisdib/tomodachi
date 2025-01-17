package user

// User defines a registered user
type User struct {
	ID         string `pg:"id" json:"id"`
	Username   string `pg:"username" json:"username"`
	Email      string `pg:"email" json:"email"`
	Birthdate  uint64 `pg:"birthdate" json:"birthdate"`
	Rating     uint8  `pg:"rating" json:"rating,omitempty"`
	IsVerified bool   `pg:"verified" json:"isVerified"`
	CreatedAt  uint64 `pg:"created_at" json:"createdAt"`
	UpdatedAt  uint64 `pg:"updated_at" json:"updatedAt,omitempty"`
}
