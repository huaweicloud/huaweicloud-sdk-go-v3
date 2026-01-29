package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigInfo struct {

	// 纳管账号列表(非跨账号场景不传递)
	Accounts *[]string `json:"accounts,omitempty"`

	// 操作
	Action *string `json:"action,omitempty"`

	// 自动转告警的开关
	Alert *bool `json:"alert,omitempty"`

	// 是否接入已纳管的全量账号
	AllAccounts *bool `json:"all_accounts,omitempty"`

	// 云产品
	Csvc *string `json:"csvc,omitempty"`

	// 云服务描述
	CsvcDisplay string `json:"csvc_display"`

	// 启用状态：0未启用；1启用
	Enable int64 `json:"enable"`

	// 新账号自动同步的开关
	NewAccountAutoAccess *bool `json:"new_account_auto_access,omitempty"`

	// 所需分区数
	Shards int64 `json:"shards"`

	// 数据源描述
	SourceDisplay string `json:"source_display"`

	// 数据源ID
	SourceId int64 `json:"source_id"`

	// 日志名称
	SourceName *string `json:"source_name,omitempty"`

	// 数据生命周期
	Ttl int64 `json:"ttl"`
}

func (o ConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigInfo struct{}"
	}

	return strings.Join([]string{"ConfigInfo", string(data)}, " ")
}
