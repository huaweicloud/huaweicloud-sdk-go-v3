package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSharedConnectionRequest Request Object
type CreateSharedConnectionRequest struct {
	Body *CreateSharedConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateSharedConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSharedConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateSharedConnectionRequest", string(data)}, " ")
}
