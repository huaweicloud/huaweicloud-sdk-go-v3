package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMasterSlavePoolRequest Request Object
type CreateMasterSlavePoolRequest struct {
	Body *CreateMasterSlavePoolRequestBody `json:"body,omitempty"`
}

func (o CreateMasterSlavePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlavePoolRequest struct{}"
	}

	return strings.Join([]string{"CreateMasterSlavePoolRequest", string(data)}, " ")
}
