package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主持人邀请开启、关闭摄像头的请求body
type RestVideoBody struct {

	// 1：关闭视频 0：开启视频
	Status int32 `json:"status" xml:"status"`
}

func (o RestVideoBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestVideoBody struct{}"
	}

	return strings.Join([]string{"RestVideoBody", string(data)}, " ")
}
