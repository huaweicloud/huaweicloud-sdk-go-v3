package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLakeFormationInstanceDefaultRequest Request Object
type UpdateLakeFormationInstanceDefaultRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`
}

func (o UpdateLakeFormationInstanceDefaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstanceDefaultRequest struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstanceDefaultRequest", string(data)}, " ")
}
