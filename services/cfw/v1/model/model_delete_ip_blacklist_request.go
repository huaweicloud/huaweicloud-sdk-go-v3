package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpBlacklistRequest Request Object
type DeleteIpBlacklistRequest struct {

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	Body *IpBlacklistDeleteDto `json:"body,omitempty"`
}

func (o DeleteIpBlacklistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpBlacklistRequest struct{}"
	}

	return strings.Join([]string{"DeleteIpBlacklistRequest", string(data)}, " ")
}
