package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarRocksDataReplicationResponse Response Object
type DeleteStarRocksDataReplicationResponse struct {

	// 删除同步任务工作ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteStarRocksDataReplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRocksDataReplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteStarRocksDataReplicationResponse", string(data)}, " ")
}
