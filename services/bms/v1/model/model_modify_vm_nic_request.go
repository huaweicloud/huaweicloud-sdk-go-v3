package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyVmNicRequest Request Object
type ModifyVmNicRequest struct {
	NicId string `json:"nic_id"`

	Body *BareMetalModifyPortRequest `json:"body,omitempty"`
}

func (o ModifyVmNicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyVmNicRequest struct{}"
	}

	return strings.Join([]string{"ModifyVmNicRequest", string(data)}, " ")
}
