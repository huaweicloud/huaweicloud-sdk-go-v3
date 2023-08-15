package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddIpGroupIpRequest Request Object
type AddIpGroupIpRequest struct {

	// IP地址组ID。
	IpGroupId string `json:"ip_group_id"`

	Body *AddIpGroupIpRequestBody `json:"body,omitempty"`
}

func (o AddIpGroupIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIpGroupIpRequest struct{}"
	}

	return strings.Join([]string{"AddIpGroupIpRequest", string(data)}, " ")
}
