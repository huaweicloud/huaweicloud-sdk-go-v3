package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowManagedAccountRequest Request Object
type ShowManagedAccountRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`
}

func (o ShowManagedAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowManagedAccountRequest", string(data)}, " ")
}
