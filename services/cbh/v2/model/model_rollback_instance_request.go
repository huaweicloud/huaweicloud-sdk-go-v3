package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackInstanceRequest Request Object
type RollbackInstanceRequest struct {
	Body *RollbackInstanceRequestBody `json:"body,omitempty"`
}

func (o RollbackInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackInstanceRequest struct{}"
	}

	return strings.Join([]string{"RollbackInstanceRequest", string(data)}, " ")
}
