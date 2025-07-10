package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueStatuesResponse Response Object
type ListIssueStatuesResponse struct {

	// 查询结果总数
	Total *int32 `json:"total,omitempty"`

	// 查询结果
	Result *[]IpdStatusVo `json:"result,omitempty"`

	// 状态码
	Status *string `json:"status,omitempty"`

	// 响应信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIssueStatuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueStatuesResponse struct{}"
	}

	return strings.Join([]string{"ListIssueStatuesResponse", string(data)}, " ")
}
