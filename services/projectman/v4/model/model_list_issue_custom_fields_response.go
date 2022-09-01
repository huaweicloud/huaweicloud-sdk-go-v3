package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListIssueCustomFieldsResponse struct {

	// 自定义字段返回数据
	Datas          *[]IssueCustomField `json:"datas,omitempty" xml:"datas"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListIssueCustomFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueCustomFieldsResponse struct{}"
	}

	return strings.Join([]string{"ListIssueCustomFieldsResponse", string(data)}, " ")
}
