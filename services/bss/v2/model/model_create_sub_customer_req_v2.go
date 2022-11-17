package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSubCustomerReqV2 struct {

	// 企业子账号挂载的组织单元。 组织单元的Party ID，通过查询企业组织结构接口的响应获得。
	PartyId string `json:"party_id"`

	// 企业子账号的显示名称，不限制特殊字符。
	DisplayName *string `json:"display_name,omitempty"`

	// 子账号关联类型：1：同一法人。 关联类型目前只能是同一法人。
	SubCustomerAssociationType int32 `json:"sub_customer_association_type"`

	// 申请的权限列表。 支持的权限项请参见下表。
	PermissionIds *[]string `json:"permission_ids,omitempty"`

	NewSubCustomer *NewCustomerV2 `json:"new_sub_customer"`
}

func (o CreateSubCustomerReqV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubCustomerReqV2 struct{}"
	}

	return strings.Join([]string{"CreateSubCustomerReqV2", string(data)}, " ")
}
