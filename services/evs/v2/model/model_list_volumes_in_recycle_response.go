package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumesInRecycleResponse Response Object
type ListVolumesInRecycleResponse struct {

	// **参数解释** 云硬盘总数。 **取值范围** 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释** 云硬盘列表。 **取值范围** 不涉及。
	Volumes        *[]RecycleBinVolume `json:"volumes,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVolumesInRecycleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesInRecycleResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesInRecycleResponse", string(data)}, " ")
}
