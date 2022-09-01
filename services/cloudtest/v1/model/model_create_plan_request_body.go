package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建计划请求体
type CreatePlanRequestBody struct {

	// 计划名称
	Name string `json:"name" xml:"name"`

	// 处理者id，不填时默认使用当前用户
	AssignedId *string `json:"assigned_id,omitempty" xml:"assigned_id"`

	// 计划下包含的用例类型，数组长度小于10个
	ServiceIdList []int32 `json:"service_id_list" xml:"service_id_list"`

	PlanCycle *PlanCycle `json:"plan_cycle" xml:"plan_cycle"`
}

func (o CreatePlanRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlanRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePlanRequestBody", string(data)}, " ")
}
