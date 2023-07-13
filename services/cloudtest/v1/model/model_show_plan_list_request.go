package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlanListRequest Request Object
type ShowPlanListRequest struct {

	// 起始偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset int64 `json:"offset"`

	// 每页显示的条目数量,最大支持200条
	Limit int64 `json:"limit"`

	// 模糊查询使用(针对测试计划名称)
	Name *string `json:"name,omitempty"`

	// 测试计划所处阶段（create,design,execute,report）
	CurrentStage *string `json:"current_stage,omitempty"`
}

func (o ShowPlanListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlanListRequest struct{}"
	}

	return strings.Join([]string{"ShowPlanListRequest", string(data)}, " ")
}
