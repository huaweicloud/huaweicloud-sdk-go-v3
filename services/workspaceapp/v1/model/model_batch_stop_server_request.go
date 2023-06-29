package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopServerRequest Request Object
type BatchStopServerRequest struct {
	Body *ServerHaltReq `json:"body,omitempty"`
}

func (o BatchStopServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServerRequest struct{}"
	}

	return strings.Join([]string{"BatchStopServerRequest", string(data)}, " ")
}
