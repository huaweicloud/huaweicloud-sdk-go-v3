package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadAccessTokenForJwtRequest Request Object
type CreateWorkloadAccessTokenForJwtRequest struct {
	Body *CreateWorkloadAccessTokenForJwtRequestBody `json:"body,omitempty"`
}

func (o CreateWorkloadAccessTokenForJwtRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenForJwtRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenForJwtRequest", string(data)}, " ")
}
