package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDataReplicationsResponse Response Object
type ListStarRocksDataReplicationsResponse struct {

	// 查询数据同步任务数。
	TotalCount *string `json:"total_count,omitempty"`

	// 数据同步任务信息。
	Replications *[]StarRocksReplicationInfo `json:"replications,omitempty"`

	// 扩展字段。
	ExtText        *string `json:"ext_text,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListStarRocksDataReplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDataReplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListStarRocksDataReplicationsResponse", string(data)}, " ")
}
