package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpIpGroupRequest Request Object
type UpdateHttpIpGroupRequest struct {

	// IP地址组id
	IpGroupId string `json:"ip_group_id"`

	Body *UpdateHttpIpGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpIpGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpIpGroupRequest", string(data)}, " ")
}
