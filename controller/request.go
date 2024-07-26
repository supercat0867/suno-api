package controller

// CreateGenerateSongTaskRequest 创建生成歌曲任务请求
type CreateGenerateSongTaskRequest struct {
	Prompt   string `json:"prompt" binding:"required"` // 歌词
	Title    string `json:"title" binding:"required"`  // 歌名
	Tags     string `json:"tags" binding:"required"`   // 风格、元素
	CallBack string `json:"callBack"`                  // 回调地址
}
