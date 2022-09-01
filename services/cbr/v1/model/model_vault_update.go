package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultUpdate struct {
	Billing *BillingUpdate `json:"billing,omitempty" xml:"billing"`

	// 存储库名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 是否支持自动挂载
	AutoBind *bool `json:"auto_bind,omitempty" xml:"auto_bind"`

	BindRules *VaultBindRules `json:"bind_rules,omitempty" xml:"bind_rules"`

	// 是否自动扩容。按需存储库支持自动扩容，包周期存储库不支持扩容。
	AutoExpand *bool `json:"auto_expand,omitempty" xml:"auto_expand"`

	// 发送smn通知开关
	SmnNotify *bool `json:"smn_notify,omitempty" xml:"smn_notify"`

	// 存储库容量阈值，存储库已用容量和总容量的百分比超过该值，若smn_notify为开，将发送相关通知。
	Threshold *int32 `json:"threshold,omitempty" xml:"threshold"`
}

func (o VaultUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultUpdate struct{}"
	}

	return strings.Join([]string{"VaultUpdate", string(data)}, " ")
}
