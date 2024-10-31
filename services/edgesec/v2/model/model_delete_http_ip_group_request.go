package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpIpGroupRequest Request Object
type DeleteHttpIpGroupRequest struct {

	// IP地址组id
	IpGroupId string `json:"ip_group_id"`
}

func (o DeleteHttpIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpIpGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpIpGroupRequest", string(data)}, " ")
}
