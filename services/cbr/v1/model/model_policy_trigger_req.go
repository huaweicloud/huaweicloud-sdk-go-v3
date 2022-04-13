package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTriggerReq struct {
	Properties *PolicyTriggerPropertiesReq `json:"properties"`
}

func (o PolicyTriggerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerReq struct{}"
	}

	return strings.Join([]string{"PolicyTriggerReq", string(data)}, " ")
}
