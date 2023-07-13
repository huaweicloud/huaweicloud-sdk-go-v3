package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProjectModuleRequest Request Object
type DeleteProjectModuleRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 模块id
	ModuleId int32 `json:"module_id"`
}

func (o DeleteProjectModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProjectModuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteProjectModuleRequest", string(data)}, " ")
}
