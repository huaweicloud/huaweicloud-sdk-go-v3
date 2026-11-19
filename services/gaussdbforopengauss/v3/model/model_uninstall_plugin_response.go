package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallPluginResponse Response Object
type UninstallPluginResponse struct {

	// **参数解释**: 工作流ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UninstallPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallPluginResponse struct{}"
	}

	return strings.Join([]string{"UninstallPluginResponse", string(data)}, " ")
}
