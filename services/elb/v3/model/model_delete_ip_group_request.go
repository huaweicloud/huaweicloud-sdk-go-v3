package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteIpGroupRequest struct {
	// IP地址组的ID。

	IpgroupId string `json:"ipgroup_id"`
}

func (o DeleteIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupRequest", string(data)}, " ")
}
