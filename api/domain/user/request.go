package user

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Birthdate string `json:"birthdate" validate:"required"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Location  string `json:"location,omitempty"`
	Website   string `json:"website,omitempty"`
	IsActive  bool   `json:"is_active" validate:"required"`
	CreatedAt string `json:"created_at" validate:"required"`
}
