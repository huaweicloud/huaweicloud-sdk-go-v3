package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAnAppV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *AppCreate `json:"body,omitempty"`
}

func (o CreateAnAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnAppV2Request struct{}"
	}

	return strings.Join([]string{"CreateAnAppV2Request", string(data)}, " ")
}
