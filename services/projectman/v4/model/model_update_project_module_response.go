package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectModuleResponse Response Object
type UpdateProjectModuleResponse struct {

	// 模块描述
	Description *string `json:"description,omitempty"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty"`

	// 模块id
	ModuleId *int32 `json:"module_id,omitempty"`

	Owner          *ModuleOwner `json:"owner,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateProjectModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectModuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectModuleResponse", string(data)}, " ")
}
