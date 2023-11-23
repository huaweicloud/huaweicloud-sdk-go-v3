package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowElbIpGroupRequest Request Object
type ShowElbIpGroupRequest struct {

	// 实例id。
	InstanceId string `json:"instance_id"`
}

func (o ShowElbIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowElbIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowElbIpGroupRequest", string(data)}, " ")
}
