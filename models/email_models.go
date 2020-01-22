package models

// EMAIL represents an email address to be added to our mailing list.
type Email struct {
	Email  string `json:"Email" binding:"required"`
	Action string
}
