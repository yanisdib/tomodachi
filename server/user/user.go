package user

// User defines a registered user
type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Birthdate  uint64 `json:"birthdate"`
	Rating     uint8  `json:"rating,omitempty"`
	IsVerified bool   `json:"isVerified"`
	CreatedAt  uint64 `json:"createdAt"`
	UpdatedAt  uint64 `json:"updatedAt,omitempty"`
}
