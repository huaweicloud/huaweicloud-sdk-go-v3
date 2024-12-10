package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGrantsRequest Request Object
type CreateGrantsRequest struct {
	Body *GrantSecretReqBody `json:"body,omitempty"`
}

func (o CreateGrantsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGrantsRequest struct{}"
	}

	return strings.Join([]string{"CreateGrantsRequest", string(data)}, " ")
}
