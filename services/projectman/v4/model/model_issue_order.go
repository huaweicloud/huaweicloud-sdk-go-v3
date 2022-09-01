package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueOrder struct {

	// 顺序id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 顺序值
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o IssueOrder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueOrder struct{}"
	}

	return strings.Join([]string{"IssueOrder", string(data)}, " ")
}
