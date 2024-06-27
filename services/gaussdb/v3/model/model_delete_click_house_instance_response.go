package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseInstanceResponse Response Object
type DeleteClickHouseInstanceResponse struct {

	// 创建的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClickHouseInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseInstanceResponse", string(data)}, " ")
}
