package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateProjectModuleResponse struct {

	// 模块描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty" xml:"module_name"`

	// 模块id
	ModuleId *int32 `json:"module_id,omitempty" xml:"module_id"`

	Owner          *ModuleOwner `json:"owner,omitempty" xml:"owner"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateProjectModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectModuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectModuleResponse", string(data)}, " ")
}