package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrecheckResultRequest Request Object
type ShowPrecheckResultRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o ShowPrecheckResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrecheckResultRequest struct{}"
	}

	return strings.Join([]string{"ShowPrecheckResultRequest", string(data)}, " ")
}
