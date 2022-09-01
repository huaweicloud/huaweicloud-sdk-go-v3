package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClusterDetailsResponse struct {
	Cluster        *ClusterDetail `json:"cluster,omitempty" xml:"cluster"`
	HttpStatusCode int            `json:"-"`
}

func (o ListClusterDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterDetailsResponse", string(data)}, " ")
}
