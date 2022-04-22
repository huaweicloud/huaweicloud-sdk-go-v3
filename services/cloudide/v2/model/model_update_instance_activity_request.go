package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceActivityRequest struct {

	// CloudIDE实例id
	InstanceId string `json:"instance_id"`
}

func (o UpdateInstanceActivityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceActivityRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceActivityRequest", string(data)}, " ")
}
