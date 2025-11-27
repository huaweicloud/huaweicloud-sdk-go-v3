package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedClustersResponse Response Object
type ListManagedClustersResponse struct {
	Body           *[]Cluster `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListManagedClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedClustersResponse struct{}"
	}

	return strings.Join([]string{"ListManagedClustersResponse", string(data)}, " ")
}
