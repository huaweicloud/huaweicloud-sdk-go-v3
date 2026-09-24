package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyTriggerUpdateReq struct {
	Properties *PolicyTriggerPropertiesUpdateReq `json:"properties"`
}

func (o PolicyTriggerUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerUpdateReq struct{}"
	}

	return strings.Join([]string{"PolicyTriggerUpdateReq", string(data)}, " ")
}
