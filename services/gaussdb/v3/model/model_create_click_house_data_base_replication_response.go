package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClickHouseDataBaseReplicationResponse Response Object
type CreateClickHouseDataBaseReplicationResponse struct {

	// 创建的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClickHouseDataBaseReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClickHouseDataBaseReplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateClickHouseDataBaseReplicationResponse", string(data)}, " ")
}
