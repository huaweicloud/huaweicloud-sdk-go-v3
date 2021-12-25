package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPlanJournalsResponse struct {
	// 项目下查询测试计划列表返回结构

	Body           *[]ShowPlanJournalsResponseBody `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowPlanJournalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlanJournalsResponse struct{}"
	}

	return strings.Join([]string{"ShowPlanJournalsResponse", string(data)}, " ")
}
