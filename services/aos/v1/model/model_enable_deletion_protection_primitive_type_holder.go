package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnableDeletionProtectionPrimitiveTypeHolder struct {

	// 删除保护的标识位，如果不传默认为false，即默认不开启资源栈删除保护（删除保护开启后资源栈不允许被删除）
	EnableDeletionProtection *bool `json:"enable_deletion_protection,omitempty"`
}

func (o EnableDeletionProtectionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDeletionProtectionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"EnableDeletionProtectionPrimitiveTypeHolder", string(data)}, " ")
}
