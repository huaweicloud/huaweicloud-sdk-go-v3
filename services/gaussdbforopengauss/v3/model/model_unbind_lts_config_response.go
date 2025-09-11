package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindLtsConfigResponse Response Object
type UnbindLtsConfigResponse struct {

	// **参数解释**: 解除LTS日志流任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnbindLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"UnbindLtsConfigResponse", string(data)}, " ")
}
