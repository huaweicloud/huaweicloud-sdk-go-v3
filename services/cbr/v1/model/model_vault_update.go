package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultUpdate struct {
	Billing *BillingUpdate `json:"billing,omitempty"`

	// 存储库名称
	Name *string `json:"name,omitempty"`

	// 是否支持自动挂载
	AutoBind *bool `json:"auto_bind,omitempty"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty"`

	// 是否自动扩容。按需存储库支持自动扩容，包周期存储库不支持扩容。
	AutoExpand *bool `json:"auto_expand,omitempty"`

	// 发送smn通知开关
	SmnNotify *bool `json:"smn_notify,omitempty"`

	// 存储库容量阈值，存储库已用容量和总容量的百分比超过该值，若smn_notify为开，将发送相关通知。
	Threshold *int32 `json:"threshold,omitempty"`

	// 用于标识当前存储库是否已锁定
	Locked *bool `json:"locked,omitempty"`
}

func (o VaultUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultUpdate struct{}"
	}

	return strings.Join([]string{"VaultUpdate", string(data)}, " ")
}
