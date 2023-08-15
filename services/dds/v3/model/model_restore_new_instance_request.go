package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreNewInstanceRequest Request Object
type RestoreNewInstanceRequest struct {
	Body *RestoreNewInstanceRequestBody `json:"body,omitempty"`
}

func (o RestoreNewInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceRequest", string(data)}, " ")
}
