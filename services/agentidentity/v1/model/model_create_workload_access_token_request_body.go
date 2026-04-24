package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWorkloadAccessTokenRequestBody struct {

	// The unique identifier for the registered workload
	WorkloadName string `json:"workload_name"`
}

func (o CreateWorkloadAccessTokenRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkloadAccessTokenRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWorkloadAccessTokenRequestBody", string(data)}, " ")
}
