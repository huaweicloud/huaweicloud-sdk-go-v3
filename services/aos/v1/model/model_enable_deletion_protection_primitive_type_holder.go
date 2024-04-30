package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnableDeletionProtectionPrimitiveTypeHolder struct {

	// 删除保护的标识位，如果不传默认为false，即默认不开启资源栈删除保护（删除保护开启后资源栈不允许被删除）  *在UpdateStack API中，如果该参数未在RequestBody中给予，则不会对资源栈的删除保护属性进行更新*
	EnableDeletionProtection *bool `json:"enable_deletion_protection,omitempty"`
}

func (o EnableDeletionProtectionPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDeletionProtectionPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"EnableDeletionProtectionPrimitiveTypeHolder", string(data)}, " ")
}
