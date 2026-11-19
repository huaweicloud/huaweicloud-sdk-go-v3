package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallPluginResponse Response Object
type InstallPluginResponse struct {

	// **参数解释**: 工作流ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InstallPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallPluginResponse struct{}"
	}

	return strings.Join([]string{"InstallPluginResponse", string(data)}, " ")
}
