package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueBySnapIdsResponse Response Object
type ListIssueBySnapIdsResponse struct {

	// 请求返回的结果信息。
	Message *string `json:"message,omitempty"`

	// 请求状态码。
	Code *string `json:"code,omitempty"`

	// 快照对应的工作项信息。
	Issues         *[]IssueVo `json:"issues,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListIssueBySnapIdsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueBySnapIdsResponse struct{}"
	}

	return strings.Join([]string{"ListIssueBySnapIdsResponse", string(data)}, " ")
}
