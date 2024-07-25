package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MongPageRequest struct {

	// **参数解释：**  结束时间。系统以数据实例的最后修改时间作为查询条件，您定义的开始时间和结束时间作为时间范围进行查询。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EndLastModifiedTime *string `json:"endLastModifiedTime,omitempty"`

	// **参数解释：**  数据实例ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  版本号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  关系实体源端ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceId *string `json:"sourceId,omitempty"`

	// **参数解释：**  关系实体源端系统版本。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceRdmVersion *int32 `json:"sourceRdmVersion,omitempty"`

	// **参数解释：**  开始时间。系统以数据实例的最后修改时间作为查询条件，您定义的开始时间和结束时间作为时间范围进行查询。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	StartLastModifiedTime *string `json:"startLastModifiedTime,omitempty"`

	// **参数解释：**  关系实体目标端ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetId *string `json:"targetId,omitempty"`

	// **参数解释：**  关系实体目标端系统版本。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetRdmVersion *int32 `json:"targetRdmVersion,omitempty"`

	// **参数解释：**  单边不确定关系的目标端类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetType *string `json:"targetType,omitempty"`
}

func (o MongPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MongPageRequest struct{}"
	}

	return strings.Join([]string{"MongPageRequest", string(data)}, " ")
}
