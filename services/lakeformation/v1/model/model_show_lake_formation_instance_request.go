package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLakeFormationInstanceRequest Request Object
type ShowLakeFormationInstanceRequest struct {

	// LakeFormation实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowLakeFormationInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLakeFormationInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowLakeFormationInstanceRequest", string(data)}, " ")
}
