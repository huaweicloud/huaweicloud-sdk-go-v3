package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBalanceStatusRequest Request Object
type ShowBalanceStatusRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowBalanceStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBalanceStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowBalanceStatusRequest", string(data)}, " ")
}
