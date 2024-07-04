package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVhostRequest Request Object
type CreateVhostRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateVhostBody `json:"body,omitempty"`
}

func (o CreateVhostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVhostRequest struct{}"
	}

	return strings.Join([]string{"CreateVhostRequest", string(data)}, " ")
}
