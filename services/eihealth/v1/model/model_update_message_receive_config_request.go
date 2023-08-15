package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageReceiveConfigRequest Request Object
type UpdateMessageReceiveConfigRequest struct {
	Body *SetMessageReceiveConfigReq `json:"body,omitempty"`
}

func (o UpdateMessageReceiveConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageReceiveConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateMessageReceiveConfigRequest", string(data)}, " ")
}
