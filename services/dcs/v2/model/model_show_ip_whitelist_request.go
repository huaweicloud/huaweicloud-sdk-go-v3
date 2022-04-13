package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowIpWhitelistRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowIpWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpWhitelistRequest struct{}"
	}

	return strings.Join([]string{"ShowIpWhitelistRequest", string(data)}, " ")
}
