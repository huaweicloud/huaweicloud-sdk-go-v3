package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTriggerResp struct {
	Properties *PolicyTriggerPropertiesResp `json:"properties"`
}

func (o PolicyTriggerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerResp struct{}"
	}

	return strings.Join([]string{"PolicyTriggerResp", string(data)}, " ")
}
