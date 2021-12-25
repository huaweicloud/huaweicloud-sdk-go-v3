package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineStageResp struct {
	// 阶段名称

	DisplayName *string `json:"display_name,omitempty"`
	// 阶段状态

	Status *string `json:"status,omitempty"`
	// 阶段执行结果

	Result *string `json:"result,omitempty"`
}

func (o PipelineStageResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineStageResp struct{}"
	}

	return strings.Join([]string{"PipelineStageResp", string(data)}, " ")
}
