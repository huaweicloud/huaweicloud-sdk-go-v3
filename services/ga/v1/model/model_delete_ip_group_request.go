package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpGroupRequest Request Object
type DeleteIpGroupRequest struct {

	// IP地址组ID。
	IpGroupId string `json:"ip_group_id"`
}

func (o DeleteIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupRequest", string(data)}, " ")
}
