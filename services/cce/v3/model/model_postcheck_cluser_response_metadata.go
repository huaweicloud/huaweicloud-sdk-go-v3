package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostcheckCluserResponseMetadata **参数解释：** 升级后确认元数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type PostcheckCluserResponseMetadata struct {

	// **参数解释：** 任务ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Uid *string `json:"uid,omitempty"`
}

func (o PostcheckCluserResponseMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostcheckCluserResponseMetadata struct{}"
	}

	return strings.Join([]string{"PostcheckCluserResponseMetadata", string(data)}, " ")
}
