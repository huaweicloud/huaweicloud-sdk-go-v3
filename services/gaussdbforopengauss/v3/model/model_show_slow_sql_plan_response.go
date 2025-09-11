package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowSqlPlanResponse Response Object
type ShowSlowSqlPlanResponse struct {

	// **参数解释**: 执行计划信息。 **取值范围**: 不涉及。
	GsExplain      *string `json:"gs_explain,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSlowSqlPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowSqlPlanResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowSqlPlanResponse", string(data)}, " ")
}
