package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrecheckCluserResponseMetadata **参数解释：** 升级前检查元数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type PrecheckCluserResponseMetadata struct {

	// **参数解释：** 检查任务ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Uid *string `json:"uid,omitempty"`
}

func (o PrecheckCluserResponseMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckCluserResponseMetadata struct{}"
	}

	return strings.Join([]string{"PrecheckCluserResponseMetadata", string(data)}, " ")
}
