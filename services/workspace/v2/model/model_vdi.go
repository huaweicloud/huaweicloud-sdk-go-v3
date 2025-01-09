package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Vdi struct {

	// 是否开启自动锁屏。取值为：false：表示关闭。true：表示开启。
	AutoLockEnable *bool `json:"auto_lock_enable,omitempty"`

	AutoLockOptions *AutoLockOptions `json:"auto_lock_options,omitempty"`

	// 是否开启断开后自动注销。取值为：0：表示关闭。1：表示开启。
	DisconnectLogoutEnable *int32 `json:"disconnect_logout_enable,omitempty"`

	DisconnectLogoutOptions *DisconnectLogoutOptions `json:"disconnect_logout_options,omitempty"`

	// 是否开启断开后自动注销。取值为：0：表示关闭。1：表示开启。
	DisconnectHibernateEnable *bool `json:"disconnect_hibernate_enable,omitempty"`

	DisconnectHibernateOptions *VdiDisconnectHibernateOptions `json:"disconnect_hibernate_options,omitempty"`

	// 是否开启断开后自动注销。取值为：0：表示关闭。1：表示开启。
	NoOperationHibernateEnable *bool `json:"no_operation_hibernate_enable,omitempty"`

	NoOperationHibernateOptions *VdiNoOperationHibernateOptions `json:"no_operation_hibernate_options,omitempty"`
}

func (o Vdi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vdi struct{}"
	}

	return strings.Join([]string{"Vdi", string(data)}, " ")
}
