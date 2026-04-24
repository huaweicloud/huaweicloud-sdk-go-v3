package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadIdentityName The name of the workload identity.
type WorkloadIdentityName struct {
}

func (o WorkloadIdentityName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadIdentityName struct{}"
	}

	return strings.Join([]string{"WorkloadIdentityName", string(data)}, " ")
}
