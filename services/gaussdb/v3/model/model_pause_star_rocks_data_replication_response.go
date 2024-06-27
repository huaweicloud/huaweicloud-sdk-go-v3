package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseStarRocksDataReplicationResponse Response Object
type PauseStarRocksDataReplicationResponse struct {

	// 恢复数据同步的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PauseStarRocksDataReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseStarRocksDataReplicationResponse struct{}"
	}

	return strings.Join([]string{"PauseStarRocksDataReplicationResponse", string(data)}, " ")
}
