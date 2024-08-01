package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckNoNewAccessReason struct {

	// 对访问权限检查结果的推理的描述。
	Description *string `json:"description,omitempty"`

	// 新增权限statement的sid标识符。
	StatementId *string `json:"statement_id,omitempty"`

	// 新增权限statement的index，从0开始。
	StatementIndex *int32 `json:"statement_index,omitempty"`
}

func (o CheckNoNewAccessReason) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNoNewAccessReason struct{}"
	}

	return strings.Join([]string{"CheckNoNewAccessReason", string(data)}, " ")
}
