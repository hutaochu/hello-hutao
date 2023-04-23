package req

type RegisterRequest struct {
	Name      string
	AvatarUrl string `json:"avatar_url"`
	Email     string
	Password  string
}
