package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationsResponse Response Object
type UpdateInstanceConfigurationsResponse struct {

	// 修改指定实例参数的任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 是否需要重启。  - true：是。 - false：否。
	RestartRequired *bool `json:"restart_required,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateInstanceConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConfigurationsResponse", string(data)}, " ")
}
