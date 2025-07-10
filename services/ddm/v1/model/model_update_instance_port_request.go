package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstancePortRequest Request Object
type UpdateInstancePortRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdatePortRequest `json:"body,omitempty"`
}

func (o UpdateInstancePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstancePortRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstancePortRequest", string(data)}, " ")
}
