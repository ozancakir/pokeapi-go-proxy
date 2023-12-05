package dto

type TranslateRequest struct {
	Text string  `form:"text" json:"text" binding:"required"`
	From *string `form:"from" json:"from"`
	To   *string `form:"to" json:"to"`
}
