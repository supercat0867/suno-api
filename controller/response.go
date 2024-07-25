package controller

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type SongItem struct {
	Id            string `json:"id"`
	VideoUrl      string `json:"video_url"`
	AudioUrl      string `json:"audio_url"`
	ImageUrl      string `json:"image_url"`
	ImageLargeUrl string `json:"image_large_url"`
}
