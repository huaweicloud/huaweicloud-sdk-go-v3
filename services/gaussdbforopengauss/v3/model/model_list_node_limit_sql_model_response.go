package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNodeLimitSqlModelResponse Response Object
type ListNodeLimitSqlModelResponse struct {

	// 限流SQL模板匹配信息
	NodeLimitSqlModelList *[]ListNodeLimitSqlModelResponseResult `json:"node_limit_sql_model_list,omitempty"`

	// 查询记录数。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置。
	Offset *int32 `json:"offset,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNodeLimitSqlModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeLimitSqlModelResponse struct{}"
	}

	return strings.Join([]string{"ListNodeLimitSqlModelResponse", string(data)}, " ")
}
