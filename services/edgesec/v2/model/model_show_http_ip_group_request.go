package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpIpGroupRequest Request Object
type ShowHttpIpGroupRequest struct {

	// IP地址组id
	IpGroupId string `json:"ip_group_id"`
}

func (o ShowHttpIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpIpGroupRequest", string(data)}, " ")
}
