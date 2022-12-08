package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UninstallInstanceEipRequest struct {

	// server_id
	ServerId string `json:"server_id"`

	Body *OperateEipRequestBody `json:"body,omitempty"`
}

func (o UninstallInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"UninstallInstanceEipRequest", string(data)}, " ")
}
