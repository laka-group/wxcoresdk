package wxcoresdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type SendMsgStruct struct {
	Content    string `json:"content"`
	Context    string `json:"context"`
	Icon       string `json:"icon"`
	RobotWxid  string `json:"robot_wxid"`
	Title      string `json:"title"`
	ToWxid     string `json:"to_wxid"`
	Type       int64  `json:"type"`
	Url        string `json:"url"`
	PayValue   int64  `json:"pay_value"`
	Media      string `json:"media"`
	MediaSize  int64  `json:"media_size"`
	PlayLength int64  `json:"play_length"`
	Thumb      string `json:"thumb"`
	ThumbSize  int64  `json:"thumb_size"`
}

type RespStruct struct {
	Code int64
	Msg  string
}

// send wx message
func (params *SendMsgStruct) SendMsg() (*RespStruct, error) {
	bytesData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", Config.GoCore.Url+Config.GoCore.Api.SendMsg, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	c := http.Client{}
	c.Timeout = 5 * time.Second
	httpResp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	var resp *RespStruct
	err = json.Unmarshal(respBytes, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
