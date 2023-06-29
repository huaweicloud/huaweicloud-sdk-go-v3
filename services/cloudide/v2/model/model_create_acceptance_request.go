package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAcceptanceRequest Request Object
type CreateAcceptanceRequest struct {
	Body *AcceptanceSchema `json:"body,omitempty"`
}

func (o CreateAcceptanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAcceptanceRequest struct{}"
	}

	return strings.Join([]string{"CreateAcceptanceRequest", string(data)}, " ")
}
