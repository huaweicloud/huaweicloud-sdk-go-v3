package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnhanceFullSqlsResponse Response Object
type ListEnhanceFullSqlsResponse struct {

	// **参数解释**: 总记录数量。 **取值范围**: 不涉及。
	TotalCount *int64 `json:"total_count,omitempty"`

	// **参数解释**: 全量SQL列表。
	FullSqls *[]FullSqlResult `json:"full_sqls,omitempty"`

	// **参数解释**: 最大查询记录数量。主要供前端交互控制使用。 **取值范围**: 不涉及。
	LimitCount     *int64 `json:"limit_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnhanceFullSqlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhanceFullSqlsResponse struct{}"
	}

	return strings.Join([]string{"ListEnhanceFullSqlsResponse", string(data)}, " ")
}
