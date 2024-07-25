package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllVersionsRequest Request Object
type ListAllVersionsRequest struct {

	// **参数解释：**  当前页。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1。
	CurPage *int32 `json:"curPage,omitempty"`

	// **参数解释：**  结束索引。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	EndIndex *int32 `json:"endIndex,omitempty"`

	// **参数解释：**  最大分页数。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1000。
	MaxPageSize *int32 `json:"maxPageSize,omitempty"`

	// **参数解释：**  每页大小。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1000。
	PageSize *int32 `json:"pageSize,omitempty"`

	// **参数解释：**  起始索引。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	StartIndex *int32 `json:"startIndex,omitempty"`

	// **参数解释：**  总页数。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	TotalPages *int32 `json:"totalPages,omitempty"`

	// **参数解释：**  总行数。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  0。
	TotalRows *int32 `json:"totalRows,omitempty"`

	// **参数解释：**  分页大小（路径参数）。  **约束限制：**  不涉及。  **取值范围：**  1-1000。  **默认取值：**  不涉及。
	PageSizePath int32 `json:"pageSizePath"`

	// **参数解释：**  当前页数（路径参数）。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  1。
	CurPagePath int32 `json:"curPagePath"`

	// **参数解释：**  应用唯一标识。  **约束限制：**  不涉及。  **取值范围：**  由英文字母和数字组成，且长度为32个字符。  **默认取值：**  不涉及。
	Identifier string `json:"identifier"`

	// **参数解释：**  数据模型的英文名称。  **约束限制：**  不涉及。  **取值范围：**  大写字母开头，只能包含字母、数字、\"_\"，且长度为[1-60]个字符。  **默认取值：**  不涉及。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionMasterDto `json:"body,omitempty"`
}

func (o ListAllVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListAllVersionsRequest", string(data)}, " ")
}
