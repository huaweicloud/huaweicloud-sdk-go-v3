package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectModuleResponse Response Object
type CreateProjectModuleResponse struct {

	// 模块描述
	Description *string `json:"description,omitempty"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty"`

	// 模块id
	ModuleId *int32 `json:"module_id,omitempty"`

	Owner          *ModuleOwner `json:"owner,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateProjectModuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectModuleResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectModuleResponse", string(data)}, " ")
}
