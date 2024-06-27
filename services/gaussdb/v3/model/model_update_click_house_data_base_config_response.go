package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClickHouseDataBaseConfigResponse Response Object
type UpdateClickHouseDataBaseConfigResponse struct {

	// 创建的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClickHouseDataBaseConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClickHouseDataBaseConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateClickHouseDataBaseConfigResponse", string(data)}, " ")
}
