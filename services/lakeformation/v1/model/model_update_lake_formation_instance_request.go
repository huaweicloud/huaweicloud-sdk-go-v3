package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLakeFormationInstanceRequest Request Object
type UpdateLakeFormationInstanceRequest struct {

	// LakeFormation实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateLakeFormationInstance `json:"body,omitempty"`
}

func (o UpdateLakeFormationInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLakeFormationInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateLakeFormationInstanceRequest", string(data)}, " ")
}
