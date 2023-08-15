package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueResponseV4Order 工作项优先级顺序
type IssueResponseV4Order struct {

	// 优先级顺序id
	Id *int32 `json:"id,omitempty"`

	// 优先级顺序名称
	Name *string `json:"name,omitempty"`
}

func (o IssueResponseV4Order) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueResponseV4Order struct{}"
	}

	return strings.Join([]string{"IssueResponseV4Order", string(data)}, " ")
}
