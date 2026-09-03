package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbsConnectionRequest Request Object
type CreateDbsConnectionRequest struct {
	Body *CreateDbsConnectionRequestBody `json:"body,omitempty"`
}

func (o CreateDbsConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbsConnectionRequest struct{}"
	}

	return strings.Join([]string{"CreateDbsConnectionRequest", string(data)}, " ")
}
