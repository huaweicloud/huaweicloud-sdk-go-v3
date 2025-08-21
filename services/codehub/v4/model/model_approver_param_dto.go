package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApproverParamDto struct {

	// 审核人id
	ApproverId *int32 `json:"approver_id,omitempty"`

	// 代码所有者
	CodeOwner *bool `json:"code_owner,omitempty"`

	// 是否接纳
	Accept *bool `json:"accept,omitempty"`
}

func (o ApproverParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproverParamDto struct{}"
	}

	return strings.Join([]string{"ApproverParamDto", string(data)}, " ")
}
