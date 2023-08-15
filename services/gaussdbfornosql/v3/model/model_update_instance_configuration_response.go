package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationResponse Response Object
type UpdateInstanceConfigurationResponse struct {

	// 修改实例参数的异步任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 实例是否需要重启。 - “true”需要重启。 - “false”不需要重启。
	RestartRequired *bool `json:"restart_required,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateInstanceConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationResponse", string(data)}, " ")
}
