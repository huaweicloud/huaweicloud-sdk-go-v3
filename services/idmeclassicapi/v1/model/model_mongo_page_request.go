package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MongoPageRequest struct {

	// **参数解释：**  数据实例ID，用于指定待查询历史版本的数据实例。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  开始时间，用于指定查询时间区间的起始点。系统以数据实例的最后修改时间作为查询条件。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	StartLastModifiedTime *string `json:"startLastModifiedTime,omitempty"`

	// **参数解释：**  结束时间，用于指定查询时间区间的结束点。系统以数据实例的最后修改时间作为查询条件。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	EndLastModifiedTime *string `json:"endLastModifiedTime,omitempty"`

	// **参数解释：**  系统版本号，用于指定查询特定版本的历史记录。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  关系实体源端ID，用于查询关系实体的历史版本信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceId *string `json:"sourceId,omitempty"`

	// **参数解释：**  关系实体源端系统版本，用于指定源端实例的特定版本。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SourceRdmVersion *int32 `json:"sourceRdmVersion,omitempty"`

	// **参数解释：**  关系实体目标端ID，用于查询关系实体的历史版本信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetId *string `json:"targetId,omitempty"`

	// **参数解释：**  关系实体目标端系统版本，用于指定目标端实例的特定版本。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetRdmVersion *int32 `json:"targetRdmVersion,omitempty"`

	// **参数解释：**  单边不确定关系的目标端类型，用于指定关系实体的目标端模型类型。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetType *string `json:"targetType,omitempty"`
}

func (o MongoPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MongoPageRequest struct{}"
	}

	return strings.Join([]string{"MongoPageRequest", string(data)}, " ")
}
