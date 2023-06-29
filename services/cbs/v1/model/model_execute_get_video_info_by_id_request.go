package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGetVideoInfoByIdRequest Request Object
type ExecuteGetVideoInfoByIdRequest struct {

	// 视频id
	VideoId string `json:"video_id"`
}

func (o ExecuteGetVideoInfoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGetVideoInfoByIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGetVideoInfoByIdRequest", string(data)}, " ")
}
