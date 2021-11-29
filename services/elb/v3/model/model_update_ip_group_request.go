package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateIpGroupRequest struct {
	// 待更新的IP地址组的ID。

	IpgroupId string `json:"ipgroup_id"`

	Body *UpdateIpGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequest", string(data)}, " ")
}
