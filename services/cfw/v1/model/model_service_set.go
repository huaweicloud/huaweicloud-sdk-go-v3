package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceSet struct {

	// 服务组id
	SetId *string `json:"set_id,omitempty"`

	// 服务组名称
	Name *string `json:"name,omitempty"`

	// 服务组描述
	Description *string `json:"description,omitempty"`

	// 服务组类型，0表示自定义服务组，1表示常用WEB服务，2表示常用远程登录和PING，3表示常用数据库
	ServiceSetType *int32 `json:"service_set_type,omitempty"`

	// 服务组被规则引用次数
	RefCount *int32 `json:"ref_count,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 协议列表，协议类型：TCP为6，UDP为17，ICMP为1，ICMPV6为58，ANY为-1,type为0手动类型时不能为空。
	Protocols *[]int32 `json:"protocols,omitempty"`
}

func (o ServiceSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSet struct{}"
	}

	return strings.Join([]string{"ServiceSet", string(data)}, " ")
}
