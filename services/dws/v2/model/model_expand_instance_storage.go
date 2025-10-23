package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandInstanceStorage **参数解释**： 磁盘扩容请求信息。 磁盘扩容后单节点有效存储容量（GB / 节点），集群总容量 = 单节点容量 x 节点数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ExpandInstanceStorage struct {

	// **参数解释**：   磁盘扩容后单节点有效存储容量（GB/节点）。   该容量必须大于当前单节点有效容量，小于等于集群规格支持的单节点最大容量，扩容容量为规格支持的步长倍数。  集群规格配置详情可根据 [集群规格详情](ShowClusterFlavor.xml) 查询。    **约束限制**：    不涉及。   **取值范围**：   不涉及。   **默认取值**：   不涉及。
	NewSize int32 `json:"new_size"`
}

func (o ExpandInstanceStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceStorage struct{}"
	}

	return strings.Join([]string{"ExpandInstanceStorage", string(data)}, " ")
}
