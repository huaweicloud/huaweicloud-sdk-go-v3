package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkSqlsResponse Response Object
type ListSparkSqlsResponse struct {
	PageInfo *SparkMarkerPageInfo `json:"page_info,omitempty"`

	// **参数解释**：详细的SparkSql作业列表，包含作业ID、状态、SQL内容等信息。
	Statements     *[]ListSparkSqlResItem `json:"statements,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListSparkSqlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkSqlsResponse struct{}"
	}

	return strings.Join([]string{"ListSparkSqlsResponse", string(data)}, " ")
}
