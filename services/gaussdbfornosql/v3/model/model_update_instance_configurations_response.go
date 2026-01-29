package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConfigurationsResponse Response Object
type UpdateInstanceConfigurationsResponse struct {

	// 修改实例参数的异步任务ID。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释：** 实例是否需要重启。 **取值范围：** - “true”需要重启。 - “false”不需要重启。
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
