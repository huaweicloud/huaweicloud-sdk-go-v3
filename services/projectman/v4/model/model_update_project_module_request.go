package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectModuleRequest Request Object
type UpdateProjectModuleRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 模块id
	ModuleId int32 `json:"module_id"`

	Body *UpdateProjectModuleRequestBody `json:"body,omitempty"`
}

func (o UpdateProjectModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectModuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateProjectModuleRequest", string(data)}, " ")
}
