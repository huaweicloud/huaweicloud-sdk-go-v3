package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWorkloadAccessTokenForUserIdRequestBody struct {

	// The unique identifier for the registered workload
	WorkloadName string `json:"workload_name"`

	// The ID of the user for whom to retrieve the access token
	UserId string `json:"user_id"`
}

func (o CreateWorkloadAccessTokenForUserIdRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenForUserIdRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenForUserIdRequestBody", string(data)}, " ")
}
