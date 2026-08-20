package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateRespDetail struct {

	// 关联的工作项ID，多个ID使用逗号分割。
	IssueId *string `json:"issue_id,omitempty"`

	// 失败原因。
	FailMsg *string `json:"fail_msg,omitempty"`

	// 操作类型标记位。
	OperationFlag *int32 `json:"operation_flag,omitempty"`

	// 修改日期。
	ModifiedDate *int64 `json:"modified_date,omitempty"`

	// 修改人。
	ModifiedBy *string `json:"modified_by,omitempty"`
}

func (o AssociateRespDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRespDetail struct{}"
	}

	return strings.Join([]string{"AssociateRespDetail", string(data)}, " ")
}
