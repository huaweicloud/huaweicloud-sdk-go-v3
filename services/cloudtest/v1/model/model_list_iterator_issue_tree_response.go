package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIteratorIssueTreeResponse Response Object
type ListIteratorIssueTreeResponse struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value          *[]WorkItemVo `json:"value,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListIteratorIssueTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIteratorIssueTreeResponse struct{}"
	}

	return strings.Join([]string{"ListIteratorIssueTreeResponse", string(data)}, " ")
}
