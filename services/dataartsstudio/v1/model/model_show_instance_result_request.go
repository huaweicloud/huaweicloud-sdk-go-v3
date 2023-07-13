package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceResultRequest Request Object
type ShowInstanceResultRequest struct {

	// projectId
	InstanceId string `json:"instance_id"`

	// workspace 信息
	Workspace string `json:"workspace"`
}

func (o ShowInstanceResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResultRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceResultRequest", string(data)}, " ")
}
