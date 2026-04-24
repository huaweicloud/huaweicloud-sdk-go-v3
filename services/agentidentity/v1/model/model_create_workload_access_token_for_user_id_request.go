package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadAccessTokenForUserIdRequest Request Object
type CreateWorkloadAccessTokenForUserIdRequest struct {
	Body *CreateWorkloadAccessTokenForUserIdRequestBody `json:"body,omitempty"`
}

func (o CreateWorkloadAccessTokenForUserIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenForUserIdRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenForUserIdRequest", string(data)}, " ")
}
