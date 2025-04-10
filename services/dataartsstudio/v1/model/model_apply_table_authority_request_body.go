package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyTableAuthorityRequestBody 获取表审批人列表参数
type ApplyTableAuthorityRequestBody struct {
	Approver *TableApprover `json:"approver,omitempty"`

	TableInfo *ListTableApproversRequestBody `json:"table_info,omitempty"`

	// 申请人列表
	Proposers *[]TableProposers `json:"proposers,omitempty"`

	// 申请理由
	Reason *string `json:"reason,omitempty"`
}

func (o ApplyTableAuthorityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyTableAuthorityRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyTableAuthorityRequestBody", string(data)}, " ")
}
