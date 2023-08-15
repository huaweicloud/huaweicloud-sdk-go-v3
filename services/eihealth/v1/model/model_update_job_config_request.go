package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobConfigRequest Request Object
type UpdateJobConfigRequest struct {
	Body *UpdateJobConfigReq `json:"body,omitempty"`
}

func (o UpdateJobConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobConfigRequest", string(data)}, " ")
}
