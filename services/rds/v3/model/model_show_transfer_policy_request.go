package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransferPolicyRequest Request Object
type ShowTransferPolicyRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o ShowTransferPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransferPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowTransferPolicyRequest", string(data)}, " ")
}
