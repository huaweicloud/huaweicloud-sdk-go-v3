package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEntityConfigurationResponse struct {

	// 修改实例参数的异步任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 参数修改涉及范围（实例，组，节点）否需要重启。 - false不需要重启 - true需要重启。
	RestartRequired *bool `json:"restart_required,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdateEntityConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEntityConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateEntityConfigurationResponse", string(data)}, " ")
}
