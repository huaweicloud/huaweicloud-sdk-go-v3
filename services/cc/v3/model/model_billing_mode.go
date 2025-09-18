package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BillingMode 计费方式。
type BillingMode struct {
	BillingMode *BillingModeEnum `json:"billing_mode"`
}

func (o BillingMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillingMode struct{}"
	}

	return strings.Join([]string{"BillingMode", string(data)}, " ")
}
