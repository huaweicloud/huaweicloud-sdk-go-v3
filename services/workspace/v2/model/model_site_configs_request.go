package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SiteConfigsRequest struct {

	// 开通服务资源使用的可用分区
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	NetworkConfig *NetworkConfigReq `json:"network_config"`

	AccessConfig *AccessConfigReq `json:"access_config"`
}

func (o SiteConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteConfigsRequest struct{}"
	}

	return strings.Join([]string{"SiteConfigsRequest", string(data)}, " ")
}
