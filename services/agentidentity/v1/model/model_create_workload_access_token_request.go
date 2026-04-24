package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadAccessTokenRequest Request Object
type CreateWorkloadAccessTokenRequest struct {
	Body *CreateWorkloadAccessTokenRequestBody `json:"body,omitempty"`
}

func (o CreateWorkloadAccessTokenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenRequest", string(data)}, " ")
}
