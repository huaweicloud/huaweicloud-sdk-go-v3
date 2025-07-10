package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowManagedAccountTemplateResponse Response Object
type ShowManagedAccountTemplateResponse struct {

	// 管理纳管账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 纳管账号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 纳管账号名称。
	AccountName *string `json:"account_name,omitempty"`

	// 模板ID。
	BlueprintProductId *string `json:"blueprint_product_id,omitempty"`

	// 模板名称。
	BlueprintProductName *string `json:"blueprint_product_name,omitempty"`

	// 模板版本。
	BlueprintProductVersion *string `json:"blueprint_product_version,omitempty"`

	// 模板部署状态，包括失败, 完成, 进行中。
	BlueprintStatus *string `json:"blueprint_status,omitempty"`

	// 模板实现类型。
	BlueprintProductImplType *string `json:"blueprint_product_impl_type,omitempty"`

	// 模板详情
	BlueprintProductImplDetail *string `json:"blueprint_product_impl_detail,omitempty"`
	HttpStatusCode             int     `json:"-"`
}

func (o ShowManagedAccountTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedAccountTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowManagedAccountTemplateResponse", string(data)}, " ")
}
