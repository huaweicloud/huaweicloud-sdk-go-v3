package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveIpGroupIpRequest Request Object
type RemoveIpGroupIpRequest struct {

	// IP地址组ID。
	IpGroupId string `json:"ip_group_id"`

	Body *RemoveIpGroupIpRequestBody `json:"body,omitempty"`
}

func (o RemoveIpGroupIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIpGroupIpRequest struct{}"
	}

	return strings.Join([]string{"RemoveIpGroupIpRequest", string(data)}, " ")
}
