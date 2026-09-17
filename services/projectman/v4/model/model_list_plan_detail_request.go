package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlanDetailRequest Request Object
type ListPlanDetailRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// **参数解释**： 发布/迭代唯一ID。可以通过[发布/迭代计划列表查询](ListPlan.xml)接口获取，响应消息体中的**id**字段的值就是发布/迭代ID。 **约束限制**： 长度为18-19位的数字字符串。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PlanId string `json:"plan_id"`
}

func (o ListPlanDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlanDetailRequest struct{}"
	}

	return strings.Join([]string{"ListPlanDetailRequest", string(data)}, " ")
}
