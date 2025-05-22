package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterStorageExpandRangeResponse Response Object
type ShowClusterStorageExpandRangeResponse struct {

	// **参数解释**： 扩容后单节点磁盘最小值（单位GB）。 **取值范围**： 不涉及。
	MinSize *int32 `json:"min_size,omitempty"`

	// **参数解释**： 扩容后单节点磁盘最大值（单位GB）。 **取值范围**： 不涉及。
	MaxSize *int32 `json:"max_size,omitempty"`

	// **参数解释**： 磁盘当前值（单位GB）。 **取值范围**： 不涉及。
	CurrentSize *int32 `json:"current_size,omitempty"`

	// **参数解释**： 磁盘步长大小（单位GB）。比如当前单节点磁盘20GB，步长为20，则扩容后单节点磁盘大小至少为40GB。 **取值范围**： 大于等于10。
	Step           *int32 `json:"step,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowClusterStorageExpandRangeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterStorageExpandRangeResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterStorageExpandRangeResponse", string(data)}, " ")
}
