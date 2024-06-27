package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetVo struct {

	// 域名组id
	SetId *string `json:"set_id,omitempty"`

	// 域名组名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 引用次数
	RefCount *int32 `json:"ref_count,omitempty"`

	// 域名组类型，0表示应用域名组，1表示网络域名组
	DomainSetType *int32 `json:"domain_set_type,omitempty"`

	// 配置状态
	ConfigStatus *int32 `json:"config_status,omitempty"`

	// 异常信息
	Message *string `json:"message,omitempty"`

	// 使用规则id列表
	Rules *[]UseRuleVo `json:"rules,omitempty"`
}

func (o DomainSetVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSetVo struct{}"
	}

	return strings.Join([]string{"DomainSetVo", string(data)}, " ")
}
