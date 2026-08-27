package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveModelConfigRequest Request Object
type RemoveModelConfigRequest struct {
	Body *RemoveModelConfigReq `json:"body,omitempty"`
}

func (o RemoveModelConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveModelConfigRequest struct{}"
	}

	return strings.Join([]string{"RemoveModelConfigRequest", string(data)}, " ")
}
