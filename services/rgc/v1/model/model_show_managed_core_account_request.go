package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowManagedCoreAccountRequest Request Object
type ShowManagedCoreAccountRequest struct {

	// 纳管账号类型。类型包括LOGGING，SECURITY和PRIMARY。
	AccountType string `json:"account_type"`
}

func (o ShowManagedCoreAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedCoreAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowManagedCoreAccountRequest", string(data)}, " ")
}
