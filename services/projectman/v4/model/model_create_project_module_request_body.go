package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProjectModuleRequestBody struct {

	// 模块描述
	Description *string `json:"description,omitempty"`

	// 模块名称
	ModuleName string `json:"module_name"`

	// 父模块id
	ParentModuleId *int32 `json:"parent_module_id,omitempty"`

	Owner *UserRequest `json:"owner"`
}

func (o CreateProjectModuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectModuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProjectModuleRequestBody", string(data)}, " ")
}
