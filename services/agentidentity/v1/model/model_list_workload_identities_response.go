package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadIdentitiesResponse Response Object
type ListWorkloadIdentitiesResponse struct {
	WorkloadIdentities *[]WorkloadIdentitySummary `json:"workload_identities,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListWorkloadIdentitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadIdentitiesResponse struct{}"
	}

	return strings.Join([]string{"ListWorkloadIdentitiesResponse", string(data)}, " ")
}
