package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseDataBaseConfigResponse Response Object
type DeleteClickHouseDataBaseConfigResponse struct {

	// 创建的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClickHouseDataBaseConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseDataBaseConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseDataBaseConfigResponse", string(data)}, " ")
}
