package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadAccessTokenResponse Response Object
type CreateWorkloadAccessTokenResponse struct {

	// An opaque token representing the identity of both the workload and the user (or just the workload if not acting on behalf of a user)
	WorkloadAccessToken *string `json:"workload_access_token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateWorkloadAccessTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenResponse", string(data)}, " ")
}
