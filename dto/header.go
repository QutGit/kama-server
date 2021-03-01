package dto

// Header Header
type Header struct {
	UID string `json:"uid" binding:"required"`
}
