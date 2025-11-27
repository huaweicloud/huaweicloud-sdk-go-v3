package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ManagedFieldsEntry 由API Server自动维护的字段所有权记录，用于Server-Side Apply的冲突检测与字段级并发管理。
type ManagedFieldsEntry struct {

	// 管理者的名称
	Manager *string `json:"manager,omitempty"`

	// 记录导致此条目创建的操作类型，只能是 Apply 或 Update 两种操作类型
	Operation *string `json:"operation,omitempty"`

	// 该管理者定义字段时所依据的资源 API 版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	// 此管理条目被创建或最后一次更新的时间戳
	Time *string `json:"time,omitempty"`

	// 固定为 \"FieldsV1\"，标记字段结构格式
	FieldsType *string `json:"fieldsType,omitempty"`

	// 用于存储实际被管理的字段信息
	FieldsV1 *interface{} `json:"fieldsV1,omitempty"`
}

func (o ManagedFieldsEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedFieldsEntry struct{}"
	}

	return strings.Join([]string{"ManagedFieldsEntry", string(data)}, " ")
}
