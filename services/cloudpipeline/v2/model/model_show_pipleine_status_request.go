package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPipleineStatusRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`
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
