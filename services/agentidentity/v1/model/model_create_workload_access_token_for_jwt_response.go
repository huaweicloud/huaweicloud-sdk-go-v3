package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkloadAccessTokenForJwtResponse Response Object
type CreateWorkloadAccessTokenForJwtResponse struct {

	// An opaque token representing the identity of both the workload and the user (or just the workload if not acting on behalf of a user)
	WorkloadAccessToken *string `json:"workload_access_token,omitempty"`

	// The date and time on which the workload access token expire.
	Expiration     *sdktime.SdkTime `json:"expiration,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateWorkloadAccessTokenForJwtResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenForJwtResponse struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenForJwtResponse", string(data)}, " ")
}
