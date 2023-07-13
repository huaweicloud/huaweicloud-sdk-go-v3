package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlansRequest Request Object
type ShowPlansRequest struct {

	// 模糊查询使用(针对测试计划名称)
	Name *string `json:"name,omitempty"`

	// 测试计划所处阶段（create,design,execute,report）
	CurrentStage *string `json:"current_stage,omitempty"`

	// 页号，取值范围为1-20000
	Offset int32 `json:"offset"`

	// 每页显示的条目数量，取值范围为1-200
	Limit int32 `json:"limit"`
}

func (o ShowPlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlansRequest struct{}"
	}

	return strings.Join([]string{"ShowPlansRequest", string(data)}, " ")
}
