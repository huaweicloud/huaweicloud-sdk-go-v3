package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResolutionInstances **参数解释：** DNS批量查询接口支持响应的维度。 **取值范围：** - dns_public_zone_id：公网域名ID - dns_public_recordset_id：公网记录集ID，需与dns_public_zone_id组合使用 - dns_private_zone_id：内网域名ID - vpc_id：VPC ID，需与dns_private_zone_id组合使用
type ResolutionInstances struct {
	DnsPublicZoneId *ZoneResolutionInstances `json:"dns_public_zone_id,omitempty"`

	DnsPublicRecordsetId *RsetResolutionInstances `json:"dns_public_recordset_id,omitempty"`

	DnsPrivateZoneId *ZoneResolutionInstances `json:"dns_private_zone_id,omitempty"`

	VpcId *VpcResolutionInstances `json:"vpc_id,omitempty"`
}

func (o ResolutionInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolutionInstances struct{}"
	}

	return strings.Join([]string{"ResolutionInstances", string(data)}, " ")
}
