package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceGroupVo struct {

	// 服务组名称
	Name *string `json:"name,omitempty"`

	// 协议列表
	Protocols *[]int32 `json:"protocols,omitempty"`

	// 服务组类型，0表示自定义服务组，1表示常用WEB服务，2表示常用远程登录和PING，3表示常用数据库
	ServiceSetType *int32 `json:"service_set_type,omitempty"`

	// 服务组ID
	SetId *string `json:"set_id,omitempty"`
}

func (o ServiceGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceGroupVo struct{}"
	}

	return strings.Join([]string{"ServiceGroupVo", string(data)}, " ")
}
