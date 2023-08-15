package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectionRequest Request Object
type CreateConnectionRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	Body *ConnectionInfo `json:"body,omitempty"`
}

func (o CreateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequest", string(data)}, " ")
}
