package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLakeFormationInstanceRequest Request Object
type DeleteLakeFormationInstanceRequest struct {

	// 是否放入回收站
	ToRecycleBin bool `json:"to_recycle_bin"`

	// LakeFormation实例ID
	InstanceId string `json:"instance_id"`
}

func (o DeleteLakeFormationInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLakeFormationInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteLakeFormationInstanceRequest", string(data)}, " ")
}
