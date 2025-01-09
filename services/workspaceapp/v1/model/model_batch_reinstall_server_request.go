package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchReinstallServerRequest Request Object
type BatchReinstallServerRequest struct {
	Body *BatchReinstallServerReq `json:"body,omitempty"`
}

func (o BatchReinstallServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchReinstallServerRequest struct{}"
	}

	return strings.Join([]string{"BatchReinstallServerRequest", string(data)}, " ")
}
