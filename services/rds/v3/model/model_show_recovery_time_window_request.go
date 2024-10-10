package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecoveryTimeWindowRequest Request Object
type ShowRecoveryTimeWindowRequest struct {

	// 操作的实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowRecoveryTimeWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecoveryTimeWindowRequest struct{}"
	}

	return strings.Join([]string{"ShowRecoveryTimeWindowRequest", string(data)}, " ")
}
