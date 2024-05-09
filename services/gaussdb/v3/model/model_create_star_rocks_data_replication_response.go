package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStarRocksDataReplicationResponse Response Object
type CreateStarRocksDataReplicationResponse struct {

	// 创建数据同步的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStarRocksDataReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStarRocksDataReplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateStarRocksDataReplicationResponse", string(data)}, " ")
}
