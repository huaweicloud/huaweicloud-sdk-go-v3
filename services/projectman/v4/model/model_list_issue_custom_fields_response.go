package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueCustomFieldsResponse Response Object
type ListIssueCustomFieldsResponse struct {

	// 自定义字段返回数据
	Datas          *[]IssueCustomField `json:"datas,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListIssueCustomFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueCustomFieldsResponse struct{}"
	}

	return strings.Join([]string{"ListIssueCustomFieldsResponse", string(data)}, " ")
}
