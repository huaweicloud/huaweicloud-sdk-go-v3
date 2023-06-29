package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipleineStatusRequest Request Object
type ShowPipleineStatusRequest struct {

	// 要获取状态的流水线ID
	PipelineId string `json:"pipeline_id"`

	// 要获取状态的执行ID
	BuildId *string `json:"build_id,omitempty"`
}

func (o ShowPipleineStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipleineStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowPipleineStatusRequest", string(data)}, " ")
}
