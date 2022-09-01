package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectModule struct {

	// 模块id
	ModuleId *int32 `json:"module_id,omitempty" xml:"module_id"`

	// 模块名称
	ModuleName *string `json:"module_name,omitempty" xml:"module_name"`

	Owner *ModuleOwner `json:"owner,omitempty" xml:"owner"`

	// 模块层级
	Deepth *int32 `json:"deepth,omitempty" xml:"deepth"`

	// 是否是父级，true 父模块， false 非父模块
	IsParent *bool `json:"is_parent,omitempty" xml:"is_parent"`

	// 子模块信息
	Children *[]ProjectChildModule `json:"children,omitempty" xml:"children"`
}

func (o ProjectModule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectModule struct{}"
	}

	return strings.Join([]string{"ProjectModule", string(data)}, " ")
}
