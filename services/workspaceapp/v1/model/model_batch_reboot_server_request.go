package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootServerRequest Request Object
type BatchRebootServerRequest struct {
	Body *ServerHaltReq `json:"body,omitempty"`
}

func (o BatchRebootServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServerRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootServerRequest", string(data)}, " ")
}
