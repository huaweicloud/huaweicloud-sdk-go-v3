package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowManagedAccountTemplateRequest Request Object
type ShowManagedAccountTemplateRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`
}

func (o ShowManagedAccountTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedAccountTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowManagedAccountTemplateRequest", string(data)}, " ")
}
