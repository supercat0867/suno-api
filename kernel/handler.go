package kernel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetSunoClient 获取sessionId
func (c *Client) GetSunoClient() (*GetSunoClientResponse, error) {
	url := "https://clerk.suno.com/v1/client?_clerk_js_version=4.73.4"
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response GetSunoClientResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, err
}

// GetJwt 获取jwt
func (c *Client) GetJwt() string {
	url := fmt.Sprintf("https://clerk.suno.com/v1/client/sessions/%s/tokens?_clerk_js_version=4.73.4", c.SessionID)
	resp, err := c.client.Post(url, "application/json", nil)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	var response GetJwtResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ""
	}
	return response.Jwt
}

// GenerateSong 生成歌曲
func (c *Client) GenerateSong(prompt, tags, title string) (*GenerateSongResponse, error) {
	url := "https://studio-api.suno.ai/api/generate/v2/"
	reqBody := map[string]interface{}{
		"mv":     "chirp-v3-5",
		"prompt": prompt,
		"tags":   tags,
		"title":  title,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.GetJwt()))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response GenerateSongResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetSongInfo 获取音乐信息
func (c *Client) GetSongInfo(songId string) (*GetSongInfoResponse, error) {
	url := fmt.Sprintf("https://studio-api.suno.ai/api/feed/v2?ids=%s", songId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.GetJwt()))
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var response GetSongInfoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
