package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMsdtcLocalHostRequest Request Object
type ShowMsdtcLocalHostRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowMsdtcLocalHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMsdtcLocalHostRequest struct{}"
	}

	return strings.Join([]string{"ShowMsdtcLocalHostRequest", string(data)}, " ")
}
