package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLakeFormationInstanceScale 变更实例规格
type UpdateLakeFormationInstanceScale struct {

	// 规格列表
	Specs *[]CreateSpec `json:"specs,omitempty"`
}

func (o UpdateLakeFormationInstanceScale) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstanceScale struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstanceScale", string(data)}, " ")
}
