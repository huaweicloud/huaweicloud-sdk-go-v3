package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceSet struct {

	// 服务组id
	SetId *string `json:"set_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 服务组类型，0表示自定义服务组，1表示常用WEB服务，2表示常用远程登录和PING，3表示常用数据库
	ServiceSetType *int32 `json:"service_set_type,omitempty"`

	// 引用次数
	RefCount *int32 `json:"ref_count,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 协议类型列表
	Protocols *[]int32 `json:"protocols,omitempty"`
}

func (o ServiceSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSet struct{}"
	}

	return strings.Join([]string{"ServiceSet", string(data)}, " ")
}
