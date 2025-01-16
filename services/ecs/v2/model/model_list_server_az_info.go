package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListServerAzInfo struct {

	// 可用区ID
	AvailabilityZoneId string `json:"availability_zone_id"`

	// 可用区类型
	Type *string `json:"type,omitempty"`

	// 可用区模式
	Mode *string `json:"mode,omitempty"`

	// 公网边界组，网络eip类别标识，用于查找az可用的eip池
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 可用区别名
	Alias *string `json:"alias,omitempty"`

	// 可用区所属的AZGroup列表
	AzGroupIds *[]string `json:"az_group_ids,omitempty"`

	// 可用区类型对应的子类型
	Category *int32 `json:"category,omitempty"`
}

func (o ListServerAzInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerAzInfo struct{}"
	}

	return strings.Join([]string{"ListServerAzInfo", string(data)}, " ")
}
