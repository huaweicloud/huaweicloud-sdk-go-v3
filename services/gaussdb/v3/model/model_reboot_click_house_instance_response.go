package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootClickHouseInstanceResponse Response Object
type RebootClickHouseInstanceResponse struct {

	// 创建的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebootClickHouseInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootClickHouseInstanceResponse struct{}"
	}

	return strings.Join([]string{"RebootClickHouseInstanceResponse", string(data)}, " ")
}
