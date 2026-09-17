package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlanResponse Response Object
type ListPlanResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： - success：查询计划列表成功 - error：查询计划列表失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 提示信息。 **取值范围**： 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**： 计划列表，包含发布及其子迭代的完整信息。
	Result *[]PlanResponseResult `json:"result,omitempty"`

	Page           *PlanListResponsePage `json:"page,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlanResponse struct{}"
	}

	return strings.Join([]string{"ListPlanResponse", string(data)}, " ")
}
