package main

// EMAIL represents an email address to be added to our mailing list.
type EMAIL struct {
	Email  string `json:"Email" binding:"required"`
	Action string
}
