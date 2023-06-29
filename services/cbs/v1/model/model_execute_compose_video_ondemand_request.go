package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteComposeVideoOndemandRequest Request Object
type ExecuteComposeVideoOndemandRequest struct {

	// 视频id
	VideoId string `json:"video_id"`
}

func (o ExecuteComposeVideoOndemandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteComposeVideoOndemandRequest struct{}"
	}

	return strings.Join([]string{"ExecuteComposeVideoOndemandRequest", string(data)}, " ")
}
