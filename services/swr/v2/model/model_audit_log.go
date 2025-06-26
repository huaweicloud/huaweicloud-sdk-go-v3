package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditLog struct {

	// audit log ID
	Id *int64 `json:"id,omitempty"`

	// 操作(e.g., create, update, delete)
	Operation *string `json:"operation,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源名称
	Resource *string `json:"resource,omitempty"`

	// 用户ID
	Username *string `json:"username,omitempty"`

	// 操作时间
	OpTime *sdktime.SdkTime `json:"op_time,omitempty"`
}

func (o AuditLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditLog struct{}"
	}

	return strings.Join([]string{"AuditLog", string(data)}, " ")
}
