package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadIdentityUrn The URN of the workload identity.
type WorkloadIdentityUrn struct {
}

func (o WorkloadIdentityUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadIdentityUrn struct{}"
	}

	return strings.Join([]string{"WorkloadIdentityUrn", string(data)}, " ")
}
