package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VdiNoOperationHibernateOptions 不操作休眠选项
type VdiNoOperationHibernateOptions struct {

	// 策略组ID。
	NoOperationHibernateMinutes *int32 `json:"no_operation_hibernate_minutes,omitempty"`
}

func (o VdiNoOperationHibernateOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VdiNoOperationHibernateOptions struct{}"
	}

	return strings.Join([]string{"VdiNoOperationHibernateOptions", string(data)}, " ")
}
