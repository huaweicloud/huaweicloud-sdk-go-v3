package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadAccessToken An opaque token representing the identity of both the workload and the user (or just the workload if not acting on behalf of a user)
type WorkloadAccessToken struct {
}

func (o WorkloadAccessToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadAccessToken struct{}"
	}

	return strings.Join([]string{"WorkloadAccessToken", string(data)}, " ")
}
