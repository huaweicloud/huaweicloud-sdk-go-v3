package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 脚本的引用详情。
type ScriptReferenceDetail struct {

	// 引用的任务名称
	ReferenceName *string `json:"reference_name,omitempty"`

	// 引用id
	ReferenceId *string `json:"reference_id,omitempty"`

	// 引用的任务类型
	ReferenceType *string `json:"reference_type,omitempty"`
}

func (o ScriptReferenceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptReferenceDetail struct{}"
	}

	return strings.Join([]string{"ScriptReferenceDetail", string(data)}, " ")
}
