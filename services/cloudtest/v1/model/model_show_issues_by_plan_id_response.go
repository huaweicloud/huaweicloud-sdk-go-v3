package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIssuesByPlanIdResponse struct {

	// 项目下某个测试计划关联的需求列表的返回结构
	Body           *[]TestPlanIssueDetail `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowIssuesByPlanIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssuesByPlanIdResponse struct{}"
	}

	return strings.Join([]string{"ShowIssuesByPlanIdResponse", string(data)}, " ")
}
