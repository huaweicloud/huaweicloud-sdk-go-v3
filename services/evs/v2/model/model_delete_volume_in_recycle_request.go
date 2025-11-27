package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVolumeInRecycleRequest Request Object
type DeleteVolumeInRecycleRequest struct {

	// **参数解释** 云硬盘ID。 可通过[查询所有云硬盘详情](evs_04_2006.xml)获取云硬盘ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	VolumeId string `json:"volume_id"`
}

func (o DeleteVolumeInRecycleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumeInRecycleRequest struct{}"
	}

	return strings.Join([]string{"DeleteVolumeInRecycleRequest", string(data)}, " ")
}
