package controller

import (
	"github.com/gin-gonic/gin"
	"suno-api/kernel"
)

type Controller struct {
	client *kernel.Client
}

func NewController(c *kernel.Client) *Controller {
	return &Controller{
		c,
	}
}

// CreateGenerateSongTask godoc
// @Summary 提交音乐生成任务
// @Description 提交suno生成任务，返回音乐id等信息。
// @Produce json
// @Param CreateGenerateSongTaskRequest body CreateGenerateSongTaskRequest true "Request Body"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/v1/suno/createTask [post]
func (c *Controller) CreateGenerateSongTask(g *gin.Context) {
	var req CreateGenerateSongTaskRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(400, Response{Code: 400, Msg: "请求体格式错误！"})
		return
	}
	// 创建音乐生成任务
	res, err := c.client.GenerateSong(req.Prompt, req.Tags, req.Title)
	if err != nil {
		g.JSON(500, Response{Code: 500, Msg: err.Error()})
		return
	}

	resp := make([]SongItem, len(res.Clips))

	for i := range res.Clips {
		resp[i] = SongItem{
			AudioUrl:      res.Clips[i].AudioUrl,
			Id:            res.Clips[i].Id,
			ImageLargeUrl: res.Clips[i].ImageLargeUrl,
			ImageUrl:      res.Clips[i].ImageUrl,
			VideoUrl:      res.Clips[i].VideoUrl,
		}
	}

	// TODO 设置回调地址，生成结束后将结果发送到回调地址

	g.JSON(200, Response{Code: 200, Msg: "success", Data: resp})
}

// GetGenerateSongStatus godoc
// @Summary 查询音乐生成状态
// @Description 通过音乐id,查询音乐生成状态
// @Produce json
// @Param songId query string true "音乐id"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/v1/suno/getStatus [get]
func (c *Controller) GetGenerateSongStatus(g *gin.Context) {
	songId := g.Query("songId")

	songInfo, err := c.client.GetSongInfo(songId)
	if err != nil {
		g.JSON(400, Response{Code: 400, Msg: err.Error()})
		return
	}

	g.JSON(200, Response{Code: 200, Msg: "success", Data: SongItem{
		AudioUrl:      songInfo.Clips[0].AudioUrl,
		Id:            songInfo.Clips[0].Id,
		ImageLargeUrl: songInfo.Clips[0].ImageLargeUrl,
		ImageUrl:      songInfo.Clips[0].ImageUrl,
		VideoUrl:      songInfo.Clips[0].VideoUrl,
		Status:        songInfo.Clips[0].Status,
	}})
}
