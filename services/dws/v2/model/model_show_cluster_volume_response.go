package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterVolumeResponse Response Object
type ShowClusterVolumeResponse struct {

	// **参数解释**： 单节点磁盘数量。 **取值范围**： 不涉及。
	Duplicate *int32 `json:"duplicate,omitempty"`

	// **参数解释**： 节点容量详情。 **取值范围**： 不涉及。
	DiskInfoList   *[]DiskInfoResponse `json:"disk_info_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowClusterVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterVolumeResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterVolumeResponse", string(data)}, " ")
}
