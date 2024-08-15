package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollUpgradeProgress 滚动升级信息。
type RollUpgradeProgress struct {

	// 已升级分片数，非独立部署返回null。
	UpgradedDnGroupNumbers *string `json:"upgraded_dn_group_numbers,omitempty"`

	// 总分片数，非独立部署返回null。
	TotalDnGroupNumbers *string `json:"total_dn_group_numbers,omitempty"`

	// 未完成升级的az，以“,”隔开，独立部署返回null。
	NotFullyUpgradedAz *string `json:"not_fully_upgraded_az,omitempty"`

	// 已升级az，以“,”隔开，独立部署返回null。
	AlreadyUpgradedAz *string `json:"already_upgraded_az,omitempty"`

	// az描述键值对Map<String,String>。
	AzDescriptionMap map[string]string `json:"az_description_map,omitempty"`
}

func (o RollUpgradeProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollUpgradeProgress struct{}"
	}

	return strings.Join([]string{"RollUpgradeProgress", string(data)}, " ")
}
