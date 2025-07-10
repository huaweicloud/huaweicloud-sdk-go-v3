package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueFieldsResponse Response Object
type ListIssueFieldsResponse struct {
	Page *PageVo `json:"page,omitempty"`

	// 返回数据
	Result         *[]FieldVo `json:"result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListIssueFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueFieldsResponse struct{}"
	}

	return strings.Join([]string{"ListIssueFieldsResponse", string(data)}, " ")
}
