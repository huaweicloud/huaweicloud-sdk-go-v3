package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryDataRequest Request Object
type ListHistoryDataRequest struct {

	// **参数解释：**  当前页码，从1开始计数。当路径参数curPagePath未指定时生效。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1。
	CurPage *int32 `json:"curPage,omitempty"`

	// **参数解释：**  结束索引，用于指定查询结果的结束位置。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	EndIndex *int32 `json:"endIndex,omitempty"`

	// **参数解释：**  最大分页数，用于限制返回的最大页数。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1000。
	MaxPageSize *int32 `json:"maxPageSize,omitempty"`

	// **参数解释：**  分页大小，即每页返回的实例数量。当路径参数pageSizePath未指定时生效。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1000。
	PageSize *int32 `json:"pageSize,omitempty"`

	// **参数解释：**  起始索引，用于指定查询结果的起始位置。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	StartIndex *int32 `json:"startIndex,omitempty"`

	// **参数解释：**  总页数，用于指定查询的总页数。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	TotalPages *int32 `json:"totalPages,omitempty"`

	// **参数解释：**  总行数，用于指定查询的总行数。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	TotalRows *int32 `json:"totalRows,omitempty"`

	// **参数解释：**  分页大小，即每页返回的实例数量。  **约束限制：**  不涉及。  **取值范围：**  1-1000。  **默认取值：**  不涉及。
	PageSizePath int32 `json:"pageSizePath"`

	// **参数解释：**  当前页数，从1开始计数。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1。
	CurPagePath int32 `json:"curPagePath"`

	// **参数解释：**  应用的唯一标识。  - 于2023年06月01日之前创建的应用，其唯一标识为该应用的名称。 - 于2023年06月01日之后创建的应用，其唯一标识为该应用的ID。 获取方法请参见[获取运行服务清单 - ListEnvs](https://support.huaweicloud.com/api-idme/ListApps.html)。  **约束限制：**  不涉及。  **取值范围：**  - 于2023年06月01日之前创建的应用：由英文字母和数字组成，长度为1-36个字符。 - 于2023年06月01日之后创建的应用：由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	Identifier string `json:"identifier"`

	// **参数解释：**  数据模型的英文名称。  **约束限制：**  不涉及。  **取值范围：**  以大写字母开头，只能包含字母、数字、“_”，且长度为1-60个字符。  **默认取值：**  不涉及。
	ModelName string `json:"modelName"`

	Body *RdmParamVoMongoPageRequest `json:"body,omitempty"`
}

func (o ListHistoryDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryDataRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryDataRequest", string(data)}, " ")
}
