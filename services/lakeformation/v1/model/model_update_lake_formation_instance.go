package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLakeFormationInstance 修改实例信息
type UpdateLakeFormationInstance struct {

	// 实例名称
	Name string `json:"name"`

	// 实例描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateLakeFormationInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstance struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstance", string(data)}, " ")
}
