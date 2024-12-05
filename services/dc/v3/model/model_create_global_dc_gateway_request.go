package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalDcGatewayRequest Request Object
type CreateGlobalDcGatewayRequest struct {
	Body *CreateGlobalDcGatewayRequestBody `json:"body,omitempty"`
}

func (o CreateGlobalDcGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalDcGatewayRequest struct{}"
	}

	return strings.Join([]string{"CreateGlobalDcGatewayRequest", string(data)}, " ")
}
