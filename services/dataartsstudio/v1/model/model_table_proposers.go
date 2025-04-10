package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableProposers 申请人
type TableProposers struct {

	// 申请人id，用户组id
	Id *string `json:"id,omitempty"`

	// 申请人姓名
	Name *string `json:"name,omitempty"`

	// 审批人类型, 0 用户  1 用户组
	Type *int32 `json:"type,omitempty"`
}

func (o TableProposers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableProposers struct{}"
	}

	return strings.Join([]string{"TableProposers", string(data)}, " ")
}
