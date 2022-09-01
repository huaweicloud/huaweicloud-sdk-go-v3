package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTriggerReq struct {
	Properties *PolicyTriggerPropertiesReq `json:"properties" xml:"properties"`
}

func (o PolicyTriggerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerReq struct{}"
	}

	return strings.Join([]string{"PolicyTriggerReq", string(data)}, " ")
}
