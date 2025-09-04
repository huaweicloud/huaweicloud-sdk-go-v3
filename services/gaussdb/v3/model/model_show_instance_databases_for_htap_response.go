package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDatabasesForHtapResponse Response Object
type ShowInstanceDatabasesForHtapResponse struct {

	// 数据库名称。
	Databases *[]string `json:"databases,omitempty"`

	// 数据库数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 查询时间戳。
	Timestamp      *int32 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceDatabasesForHtapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDatabasesForHtapResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceDatabasesForHtapResponse", string(data)}, " ")
}
