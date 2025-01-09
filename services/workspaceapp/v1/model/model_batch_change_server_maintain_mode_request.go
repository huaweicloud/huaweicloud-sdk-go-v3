package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeServerMaintainModeRequest Request Object
type BatchChangeServerMaintainModeRequest struct {
	Body *BatchChangeMaintainServerReq `json:"body,omitempty"`
}

func (o BatchChangeServerMaintainModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeServerMaintainModeRequest struct{}"
	}

	return strings.Join([]string{"BatchChangeServerMaintainModeRequest", string(data)}, " ")
}
