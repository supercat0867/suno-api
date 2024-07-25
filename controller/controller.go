package controller

import (
	"github.com/gin-gonic/gin"
	"suno-api/kernel"
	"sync"
	"time"
)

type Controller struct {
	client *kernel.Client
}

func NewController(c *kernel.Client) *Controller {
	return &Controller{
		c,
	}
}

// GenerateSong godoc
// @Summary 生成AI音乐
// @Description 通过歌词、歌名、风格元素、生成2首AI音乐。时间较长，一般在2～6分钟出结果。
// @Produce json
// @Param GenerateSongRequest body GenerateSongRequest true "Request Body"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/v1/generate/song [post]
func (c *Controller) GenerateSong(g *gin.Context) {
	var req GenerateSongRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(200, Response{Code: 400, Msg: "请求体格式错误！"})
		return
	}

	// 创建音乐生成任务
	res, err := c.client.GenerateSong(req.Prompt, req.Tags, req.Title)
	if err != nil {
		g.JSON(200, Response{Code: 500, Msg: err.Error()})
		return
	}

	type result struct {
		index int
		item  SongItem
		err   error
	}

	time.Sleep(time.Second * 40)

	resultChan := make(chan result, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				songInfo, err := c.client.GetSongInfo(res.Clips[i].Id)
				if err != nil {
					resultChan <- result{index: i, err: err}
					return
				}

				if songInfo.Clips[0].Status == "complete" {
					resultChan <- result{
						index: i,
						item: SongItem{
							AudioUrl:      songInfo.Clips[0].AudioUrl,
							Id:            songInfo.Clips[0].Id,
							ImageLargeUrl: songInfo.Clips[0].ImageLargeUrl,
							ImageUrl:      songInfo.Clips[0].ImageUrl,
							VideoUrl:      songInfo.Clips[0].VideoUrl,
						},
					}
					return
				}
				time.Sleep(time.Second * 3)
			}
		}(i)
	}

	wg.Wait()
	close(resultChan)

	results := make([]SongItem, 2)
	for res := range resultChan {
		if res.err != nil {
			g.JSON(500, Response{Code: 500, Msg: res.err.Error()})
			return
		}
		results[res.index] = res.item
	}

	g.JSON(200, Response{Code: 200, Msg: "success", Data: results})
}
