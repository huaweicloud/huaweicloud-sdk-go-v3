package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExternalListPeerLinks struct {
	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	GlobalDcGatewayId *string `json:"global_dc_gateway_id,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	CreateOwner *string `json:"create_owner,omitempty"`

	BandwidthInfo *BandwidthInfoExternal `json:"bandwidth_info,omitempty"`

	PeerSite *PeerSiteExternal `json:"peer_site,omitempty"`

	Status *string `json:"status,omitempty"`

	Reason *string `json:"reason,omitempty"`

	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
}

func (o ExternalListPeerLinks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalListPeerLinks struct{}"
	}

	return strings.Join([]string{"ExternalListPeerLinks", string(data)}, " ")
}
