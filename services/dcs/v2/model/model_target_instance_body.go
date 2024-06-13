package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetInstanceBody 目标实例信息。
type TargetInstanceBody struct {

	// Redis实例ID（target_instance信息中必须填写）。
	Id string `json:"id"`

	// Redis实例名称(target_instance信息中填写)。
	Name *string `json:"name,omitempty"`

	// Redis密码，如果设置了密码，则必须填写。
	Password *string `json:"password,omitempty"`

	// 任务状态。
	TaskStatus *string `json:"task_status,omitempty"`

	// Redis IP地址。
	Ip *string `json:"ip,omitempty"`

	// Redis端口。
	Port *string `json:"port,omitempty"`

	// Redis实例地址。
	Addrs *string `json:"addrs,omitempty"`

	// proxy实例是否开启了多DB。
	ProxyMultiDb *bool `json:"proxy_multi_db,omitempty"`

	// Redis数据库。
	Db *string `json:"db,omitempty"`
}

func (o TargetInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetInstanceBody struct{}"
	}

	return strings.Join([]string{"TargetInstanceBody", string(data)}, " ")
}
