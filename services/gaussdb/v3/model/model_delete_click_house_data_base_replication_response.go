package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseDataBaseReplicationResponse Response Object
type DeleteClickHouseDataBaseReplicationResponse struct {

	// 创建的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClickHouseDataBaseReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseDataBaseReplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseDataBaseReplicationResponse", string(data)}, " ")
}
