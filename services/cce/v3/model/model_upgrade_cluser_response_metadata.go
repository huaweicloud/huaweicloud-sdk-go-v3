package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeCluserResponseMetadata **参数解释：** 升级任务元数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeCluserResponseMetadata struct {

	// **参数解释：** 升级任务ID，可通过调用获取集群升级任务详情API查询进展 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Uid *string `json:"uid,omitempty"`
}

func (o UpgradeCluserResponseMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeCluserResponseMetadata struct{}"
	}

	return strings.Join([]string{"UpgradeCluserResponseMetadata", string(data)}, " ")
}
