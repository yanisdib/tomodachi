package user

type User struct {
	ID         int64  `json:"id" postgres:"id"`
	Email      string `json:"email" postgres:"email"`
	Username   string `json:"username" postgres:"username"`
	Password   string `json:"password" postgres:"password"`
	Role       Role   `json:"role" postgres:"role"`
	Birthdate  string `json:"birthdate" postgres:"birthdate"`
	AvatarURL  string `json:"avatar_url,omitempty" postgres:"avatar_url"`
	Bio        string `json:"bio,omitempty" postgres:"bio"`
	Location   string `json:"location,omitempty" postgres:"location"`
	Website    string `json:"website,omitempty" postgres:"website"`
	IsActive   bool   `json:"is_active" postgres:"active"`
	IsVerified bool   `json:"is_verified" postgres:"verified"`
	CreatedAt  string `json:"created_at" postgres:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty" postgres:"updated_at"`
	DeletedAt  string `json:"deleted_at,omitempty" postgres:"deleted_at"`
}

type Role string

const (
	Default Role = "default"
	Admin   Role = "admin"
)
