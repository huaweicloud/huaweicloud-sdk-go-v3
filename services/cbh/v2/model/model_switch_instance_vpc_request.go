package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchInstanceVpcRequest Request Object
type SwitchInstanceVpcRequest struct {
	Body *SwitchVirtualPrivateCloudRequestBody `json:"body,omitempty"`
}

func (o SwitchInstanceVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchInstanceVpcRequest struct{}"
	}

	return strings.Join([]string{"SwitchInstanceVpcRequest", string(data)}, " ")
}
