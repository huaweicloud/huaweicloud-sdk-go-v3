package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceVersionRequest Request Object
type UpgradeInstanceVersionRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *OpenGaussUpgradeRequest `json:"body,omitempty"`
}

func (o UpgradeInstanceVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceVersionRequest struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceVersionRequest", string(data)}, " ")
}
