package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigResponse Response Object
type UpdateInstanceConfigResponse struct {

	// 更新配置任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigResponse", string(data)}, " ")
}
