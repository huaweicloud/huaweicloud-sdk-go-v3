package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SwitchShardResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchShardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchShardResponse struct{}"
	}

	return strings.Join([]string{"SwitchShardResponse", string(data)}, " ")
}
