package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionRequest Request Object
type CreateConnectionRequest struct {
	Body *ConnectionCreateReq `json:"body,omitempty"`
}

func (o CreateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequest", string(data)}, " ")
}
