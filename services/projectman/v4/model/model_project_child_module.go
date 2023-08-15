package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectChildModule struct {

	// 模块id
	ModuleId *int32 `json:"module_id,omitempty"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty"`

	Owner *ModuleOwner `json:"owner,omitempty"`

	// 模块层级
	Deepth *int32 `json:"deepth,omitempty"`

	// 是否是父级，true 父模块， false 非父模块
	IsParent *bool `json:"is_parent,omitempty"`

	// 父模块id
	ParentModuleId *int32 `json:"parent_module_id,omitempty"`
}

func (o ProjectChildModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectChildModule struct{}"
	}

	return strings.Join([]string{"ProjectChildModule", string(data)}, " ")
}
