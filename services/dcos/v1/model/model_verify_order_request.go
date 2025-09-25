package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VerifyOrderRequest Request Object
type VerifyOrderRequest struct {

	// 服务单号
	Number string `json:"number"`

	Body *VerifyOrderRequestBody `json:"body,omitempty"`
}

func (o VerifyOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VerifyOrderRequest struct{}"
	}

	return strings.Join([]string{"VerifyOrderRequest", string(data)}, " ")
}
