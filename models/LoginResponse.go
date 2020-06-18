package models

//LoginResponse modelo para almacenar el token del Usuario Logueado
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
