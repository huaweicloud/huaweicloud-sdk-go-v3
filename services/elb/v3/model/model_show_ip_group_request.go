package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpGroupRequest Request Object
type ShowIpGroupRequest struct {

	// IP地址组ID。
	IpgroupId string `json:"ipgroup_id"`
}

func (o ShowIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowIpGroupRequest", string(data)}, " ")
}
