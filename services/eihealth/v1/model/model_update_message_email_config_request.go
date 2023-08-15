package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMessageEmailConfigRequest Request Object
type UpdateMessageEmailConfigRequest struct {
	Body *SetMessageEmailConfigReq `json:"body,omitempty"`
}

func (o UpdateMessageEmailConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMessageEmailConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateMessageEmailConfigRequest", string(data)}, " ")
}
