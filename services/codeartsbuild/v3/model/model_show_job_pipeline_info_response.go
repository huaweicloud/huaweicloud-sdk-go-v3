package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobPipelineInfoResponse Response Object
type ShowJobPipelineInfoResponse struct {
	Result *JobPipelineInfoItems `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobPipelineInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobPipelineInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowJobPipelineInfoResponse", string(data)}, " ")
}
