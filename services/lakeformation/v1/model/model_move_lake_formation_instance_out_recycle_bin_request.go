package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveLakeFormationInstanceOutRecycleBinRequest Request Object
type MoveLakeFormationInstanceOutRecycleBinRequest struct {

	// LakeFormation实例ID
	InstanceId string `json:"instance_id"`
}

func (o MoveLakeFormationInstanceOutRecycleBinRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveLakeFormationInstanceOutRecycleBinRequest struct{}"
	}

	return strings.Join([]string{"MoveLakeFormationInstanceOutRecycleBinRequest", string(data)}, " ")
}
