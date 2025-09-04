package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryTableRequestV3 查询HTAP主实例数据库表列表请求体。
type QueryTableRequestV3 struct {

	// **参数解释**：  查询的数据库及表名称的列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	DatabaseTables *[]DatabaseTablesInfo `json:"database_tables,omitempty"`

	// **参数解释**：  需要查询数据库的源实例ID，严格匹配UUID规则。  **约束限制**：  只能由英文字母、数字组成，且长度为36个字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// **参数解释**：  已选择的数据库及表名称的列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SelectedTables *[]DatabaseTablesInfo `json:"selected_tables,omitempty"`

	// **参数解释**：  表黑白名单设置。include_tables：白名单，exclude_tables：黑名单。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	FilterType *string `json:"filter_type,omitempty"`
}

func (o QueryTableRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTableRequestV3 struct{}"
	}

	return strings.Join([]string{"QueryTableRequestV3", string(data)}, " ")
}
