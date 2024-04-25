package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SourceInstanceBody 源实例信息。
type SourceInstanceBody struct {

	// Redis实例名称(source_instance信息中填写)。
	Addrs string `json:"addrs"`

	// Redis密码，如果设置了密码，则必须填写。
	Password *string `json:"password,omitempty"`

	// 任务状态。
	TaskStatus *string `json:"task_status,omitempty"`

	// Redis实例ID。
	Id *string `json:"id,omitempty"`

	// Redis IP地址。
	Ip *string `json:"ip,omitempty"`

	// Redis端口。
	Port *string `json:"port,omitempty"`

	// Redis名称。
	Name *string `json:"name,omitempty"`

	// proxy实例是否开启了多DB。
	ProxyMultiDb *bool `json:"proxy_multi_db,omitempty"`

	// Redis数据库。
	Db *string `json:"db,omitempty"`
}

func (o SourceInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceInstanceBody struct{}"
	}

	return strings.Join([]string{"SourceInstanceBody", string(data)}, " ")
}
