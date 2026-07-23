package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDocParamDto struct {

	// **参数解释：**  实例ID，用于筛选与指定数据模型实例关联的结构化文档。 例如查询某产品实例下的所有设计文档。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：**  文档类型，用于筛选指定类型的结构化文档。  **约束限制：**  不涉及。  **取值范围：**  - directory：目录，用于组织和管理文档层级结构。 - pageDocument：Page文档，适用于富文本编辑场景，如设计说明书、技术文档等。 - boardDocument：Board文档，适用于白板协作场景，如工艺评审、方案讨论等。 - mindDocument：Mind文档，适用于思维导图场景，如产品结构分析、流程梳理等。 - drawDocument：Draw文档，适用于绘图场景，如工艺流程图、设备布局图等。  **默认取值：**  不涉及。
	Type *string `json:"type,omitempty"`
}

func (o QueryDocParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDocParamDto struct{}"
	}

	return strings.Join([]string{"QueryDocParamDto", string(data)}, " ")
}
