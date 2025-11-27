package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyDefinitionRequest Request Object
type ShowPolicyDefinitionRequest struct {

	// 策略定义id
	Policydefinitionid string `json:"policydefinitionid"`
}

func (o ShowPolicyDefinitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyDefinitionRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyDefinitionRequest", string(data)}, " ")
}
