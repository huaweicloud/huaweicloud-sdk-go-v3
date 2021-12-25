package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIssuesByPlanIdResponse struct {
	// 查询某个测试计划关联的需求列表

	Body           *[]ShowIssuesByPlanIdResponseBody `json:"body,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ShowIssuesByPlanIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssuesByPlanIdResponse struct{}"
	}

	return strings.Join([]string{"ShowIssuesByPlanIdResponse", string(data)}, " ")
}
