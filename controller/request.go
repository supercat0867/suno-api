package controller

type GenerateSongRequest struct {
	Prompt string `json:"prompt" binding:"required"` // 歌词
	Title  string `json:"title" binding:"required"`  // 歌名
	Tags   string `json:"tags" binding:"required"`   // 风格、元素
}
