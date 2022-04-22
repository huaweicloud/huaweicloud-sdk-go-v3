package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubCustomerInfoV2 struct {

	// 企业子账号的客户ID。
	Id *string `json:"id,omitempty"`

	// 企业子账号的用户名。
	Name *string `json:"name,omitempty"`

	// 企业子账号的显示名称。 不限制特殊字符。
	DisplayName *string `json:"display_name,omitempty"`

	// 子账号状态： 1：正常2：创建中3：关闭中4：已关闭101：子账号注册中102：子账号待激活
	Status *int32 `json:"status,omitempty"`

	// 子账号归属的组织单元ID。
	OrgId *string `json:"org_id,omitempty"`

	// 子账号归属的组织单元名称。  说明： 当子账号归属的组织是企业组织根节点时，本属性可能为空。
	OrgName *string `json:"org_name,omitempty"`
}

func (o SubCustomerInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubCustomerInfoV2 struct{}"
	}

	return strings.Join([]string{"SubCustomerInfoV2", string(data)}, " ")
}
