package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectSettingsInheritCfgResponse Response Object
type UpdateProjectSettingsInheritCfgResponse struct {

	// 项目继承设置列表
	Body           *[]ProjectSettingsInheritCfgDto `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o UpdateProjectSettingsInheritCfgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectSettingsInheritCfgResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectSettingsInheritCfgResponse", string(data)}, " ")
}
