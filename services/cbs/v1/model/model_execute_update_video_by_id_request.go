package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteUpdateVideoByIdRequest Request Object
type ExecuteUpdateVideoByIdRequest struct {

	// 视频id
	VideoId string `json:"video_id"`

	Body *UpdateReq `json:"body,omitempty"`
}

func (o ExecuteUpdateVideoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUpdateVideoByIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteUpdateVideoByIdRequest", string(data)}, " ")
}
