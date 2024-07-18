package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpGroupRequest Request Object
type DeleteIpGroupRequest struct {

	// 参数解释：IP地址组的ID。
	IpgroupId string `json:"ipgroup_id"`
}

func (o DeleteIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupRequest", string(data)}, " ")
}
