package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteComposeVideoRequest Request Object
type ExecuteComposeVideoRequest struct {

	// 视频id
	VideoId string `json:"video_id"`
}

func (o ExecuteComposeVideoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteComposeVideoRequest struct{}"
	}

	return strings.Join([]string{"ExecuteComposeVideoRequest", string(data)}, " ")
}
