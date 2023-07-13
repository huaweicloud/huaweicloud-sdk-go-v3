package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoResult struct {

	// 是否是活体。
	Alive bool `json:"alive"`

	// 动作列表。
	Actions []ActionsList `json:"actions"`

	// 检测出最大人脸的图片base64。
	Picture string `json:"picture"`
}

func (o VideoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoResult struct{}"
	}

	return strings.Join([]string{"VideoResult", string(data)}, " ")
}
