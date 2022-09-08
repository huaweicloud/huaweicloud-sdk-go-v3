package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProjectModuleRequestBody struct {

	// 模块描述
	Description *string `json:"description,omitempty"`

	// 模块名称
	ModuleName string `json:"module_name"`

	Owner *UserRequest `json:"owner"`
}

func (o UpdateProjectModuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectModuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProjectModuleRequestBody", string(data)}, " ")
}
