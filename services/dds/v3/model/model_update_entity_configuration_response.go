package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEntityConfigurationResponse Response Object
type UpdateEntityConfigurationResponse struct {

	// **参数解释：** 修改实例参数的异步任务ID。 **取值范围：** 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释：** 参数修改涉及范围（实例，组，节点）否需要重启。 **取值范围：** - 取值为false，不需要重启。 - 取值为true，需要重启。
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
