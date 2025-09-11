package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbEncryptInstancePeriodRequest Request Object
type CreateDbEncryptInstancePeriodRequest struct {
	Body *ChargeOrderDbssNew `json:"body,omitempty"`
}

func (o CreateDbEncryptInstancePeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbEncryptInstancePeriodRequest struct{}"
	}

	return strings.Join([]string{"CreateDbEncryptInstancePeriodRequest", string(data)}, " ")
}
