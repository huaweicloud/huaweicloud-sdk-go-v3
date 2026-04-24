package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWorkloadAccessTokenForJwtRequestBody struct {

	// The unique identifier for the registered workload
	WorkloadName string `json:"workload_name"`

	// The OAuth 2.0 token issued by the user's identity provider
	UserToken string `json:"user_token"`
}

func (o CreateWorkloadAccessTokenForJwtRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenForJwtRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenForJwtRequestBody", string(data)}, " ")
}
