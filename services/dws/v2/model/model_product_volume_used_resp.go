package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductVolumeUsedResp **参数解释**： 集群使用的规格信息。 **取值范围**： 不涉及。
type ProductVolumeUsedResp struct {

	// **参数解释**： 节点使用存储类型。 **取值范围**： HIGH：SAS盘； ULTRAHIGH：SSD云盘； COMMON：SATA盘； LOCAL_DISK：本地盘；
	VolumeType *string `json:"volume_type,omitempty"`

	// **参数解释**： 节点使用的磁盘数量信息。 **取值范围**： 不涉及。
	VolumeNum *int32 `json:"volume_num,omitempty"`

	// **参数解释**： 集群单节点的可用存储容量。 **取值范围**： 不涉及。
	Capacity *int32 `json:"capacity,omitempty"`

	// **参数解释**： 集群节点上单数据磁盘的物理存储容量。 **取值范围**： 不涉及。
	VolumeSize *int32 `json:"volume_size,omitempty"`
}

func (o ProductVolumeUsedResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductVolumeUsedResp struct{}"
	}

	return strings.Join([]string{"ProductVolumeUsedResp", string(data)}, " ")
}
