package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 模块
type IssueItemSfV4Module struct {

	// 模块id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 模块
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o IssueItemSfV4Module) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueItemSfV4Module struct{}"
	}

	return strings.Join([]string{"IssueItemSfV4Module", string(data)}, " ")
}
