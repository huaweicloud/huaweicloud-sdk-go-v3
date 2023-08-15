package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteDeleteVideoByIdRequest Request Object
type ExecuteDeleteVideoByIdRequest struct {

	// 视频id
	VideoId string `json:"video_id"`
}

func (o ExecuteDeleteVideoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteDeleteVideoByIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteDeleteVideoByIdRequest", string(data)}, " ")
}
