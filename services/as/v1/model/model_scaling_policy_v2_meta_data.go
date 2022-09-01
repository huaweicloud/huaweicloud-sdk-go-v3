package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 附件信息
type ScalingPolicyV2MetaData struct {

	// 伸缩带宽策略中带宽对应的共享类型
	MetadataBandwidthShareType *string `json:"metadata_bandwidth_share_type,omitempty" xml:"metadata_bandwidth_share_type"`

	// 伸缩带宽策略中带宽对应的EIP的ID
	MetadataEipId *string `json:"metadata_eip_id,omitempty" xml:"metadata_eip_id"`

	// 伸缩带宽策略中带宽对应的EIP地址
	MetadataEipAddress *string `json:"metadata_eip_address,omitempty" xml:"metadata_eip_address"`
}

func (o ScalingPolicyV2MetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyV2MetaData struct{}"
	}

	return strings.Join([]string{"ScalingPolicyV2MetaData", string(data)}, " ")
}
