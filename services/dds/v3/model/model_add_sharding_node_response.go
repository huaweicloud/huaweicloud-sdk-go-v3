package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddShardingNodeResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddShardingNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddShardingNodeResponse struct{}"
	}

	return strings.Join([]string{"AddShardingNodeResponse", string(data)}, " ")
}
