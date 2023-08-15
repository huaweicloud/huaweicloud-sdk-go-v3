package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnhancedConnectionRequest Request Object
type CreateEnhancedConnectionRequest struct {
	Body *CreateEnhancedConnectionsReq `json:"body,omitempty"`
}

func (o CreateEnhancedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnhancedConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateEnhancedConnectionRequest", string(data)}, " ")
}
