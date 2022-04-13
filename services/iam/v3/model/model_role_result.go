package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RoleResult struct {
	// 权限所属账号ID。

	DomainId *string `json:"domain_id,omitempty"`
	// 该参数值为fine_grained时，标识此权限为系统内置的策略。

	Flag *string `json:"flag,omitempty"`
	// 权限的中文描述信息。

	DescriptionCn *string `json:"description_cn,omitempty"`
	// 权限所在目录。

	Catalog *string `json:"catalog,omitempty"`
	// 权限名。携带在用户的token中，云服务根据该名称来判断用户是否有权限访问。

	Name string `json:"name"`
	// 权限描述信息。

	Description *string `json:"description,omitempty"`

	Links *Links `json:"links,omitempty"`
	// 权限ID。

	Id string `json:"id"`
	// 权限展示名。

	DisplayName *string `json:"display_name,omitempty"`
	// 权限的显示模式。 > - AX表示在domain层显示。 > - XA表示在project层显示。 > - AA表示在domain和project层均显示。 > - XX表示在domain和project层均不显示。 > - 自定义策略的显示模式只能为AX或者XA，不能在domain层和project层都显示（AA），或者在domain层和project层都不显示（XX）。

	Type string `json:"type"`

	Policy *RolePolicy `json:"policy"`
	// 权限更新时间。

	UpdatedTime *string `json:"updated_time,omitempty"`
	// 权限创建时间。

	CreatedTime *string `json:"created_time,omitempty"`
}

func (o RoleResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleResult struct{}"
	}

	return strings.Join([]string{"RoleResult", string(data)}, " ")
}
