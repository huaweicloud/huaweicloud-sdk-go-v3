package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlanJournalsResponse Response Object
type ShowPlanJournalsResponse struct {

	// 项目下查询测试计划操作历史返回结构
	Body           *[]TestPlanJournalList `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowPlanJournalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlanJournalsResponse struct{}"
	}

	return strings.Join([]string{"ShowPlanJournalsResponse", string(data)}, " ")
}
