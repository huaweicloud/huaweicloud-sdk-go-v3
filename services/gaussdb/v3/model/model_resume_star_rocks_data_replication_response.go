package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeStarRocksDataReplicationResponse Response Object
type ResumeStarRocksDataReplicationResponse struct {

	// 恢复数据同步的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResumeStarRocksDataReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeStarRocksDataReplicationResponse struct{}"
	}

	return strings.Join([]string{"ResumeStarRocksDataReplicationResponse", string(data)}, " ")
}
