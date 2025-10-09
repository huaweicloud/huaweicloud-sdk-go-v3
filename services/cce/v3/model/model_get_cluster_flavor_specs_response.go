package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClusterFlavorSpecsResponse Response Object
type GetClusterFlavorSpecsResponse struct {
	ClusterFlavorSpecs *ClusterFlavorSpecification `json:"clusterFlavorSpecs,omitempty"`
	HttpStatusCode     int                         `json:"-"`
}

func (o GetClusterFlavorSpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterFlavorSpecsResponse struct{}"
	}

	return strings.Join([]string{"GetClusterFlavorSpecsResponse", string(data)}, " ")
}
