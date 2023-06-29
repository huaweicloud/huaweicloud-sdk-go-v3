package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestVideoBody 主持人邀请与会者开启/关闭摄像头请求。
type RestVideoBody struct {

	// 主持人邀请与会者开启/关闭摄像头请求。 * 1：关闭视频 * 0：开启视频
	Status int32 `json:"status"`
}

func (o RestVideoBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestVideoBody struct{}"
	}

	return strings.Join([]string{"RestVideoBody", string(data)}, " ")
}
