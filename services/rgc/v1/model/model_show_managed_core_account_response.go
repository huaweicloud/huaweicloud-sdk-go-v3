package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowManagedCoreAccountResponse Response Object
type ShowManagedCoreAccountResponse struct {

	// 纳管账号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 纳管账号名称。
	AccountName *string `json:"account_name,omitempty"`

	// 核心资源映射。
	CoreResourceMappings map[string]string `json:"core_resource_mappings,omitempty"`
	HttpStatusCode       int               `json:"-"`
}

func (o ShowManagedCoreAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedCoreAccountResponse struct{}"
	}

	return strings.Join([]string{"ShowManagedCoreAccountResponse", string(data)}, " ")
}
