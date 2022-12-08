package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type InstallInstanceEipRequest struct {

	// server_id
	ServerId string `json:"server_id"`

	Body *OperateEipRequestBody `json:"body,omitempty"`
}

func (o InstallInstanceEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallInstanceEipRequest struct{}"
	}

	return strings.Join([]string{"InstallInstanceEipRequest", string(data)}, " ")
}
