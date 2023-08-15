package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCamInstanceRequest Request Object
type CreateCamInstanceRequest struct {
	Body *InstanceCreation `json:"body,omitempty"`
}

func (o CreateCamInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCamInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateCamInstanceRequest", string(data)}, " ")
}
