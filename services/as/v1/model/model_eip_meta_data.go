package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户自定义键值对
type EipMetaData struct {

	// 伸缩带宽策略中带宽对应的共享类型。
	MetadataBandwidthShareType *string `json:"metadata_bandwidth_share_type,omitempty" xml:"metadata_bandwidth_share_type"`

	// 伸缩带宽策略中带宽对应的EIP的ID。
	MetadataEipId *string `json:"metadata_eip_id,omitempty" xml:"metadata_eip_id"`

	// 伸缩带宽策略中带宽对应的EIP地址。
	MetadataeipAddress *string `json:"metadataeip_address,omitempty" xml:"metadataeip_address"`
}

func (o EipMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipMetaData struct{}"
	}

	return strings.Join([]string{"EipMetaData", string(data)}, " ")
}
