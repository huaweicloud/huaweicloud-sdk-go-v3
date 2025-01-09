package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VdiDisconnectHibernateOptions 休眠断开选项
type VdiDisconnectHibernateOptions struct {

	// 断连休眠时间。
	DisconnectHibernateMinutes *int32 `json:"disconnect_hibernate_minutes,omitempty"`
}

func (o VdiDisconnectHibernateOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VdiDisconnectHibernateOptions struct{}"
	}

	return strings.Join([]string{"VdiDisconnectHibernateOptions", string(data)}, " ")
}
