package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionInfoRequest Request Object
type CreateConnectionInfoRequest struct {
	Body *BaseConnectionInfo `json:"body,omitempty"`
}

func (o CreateConnectionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionInfoRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionInfoRequest", string(data)}, " ")
}
