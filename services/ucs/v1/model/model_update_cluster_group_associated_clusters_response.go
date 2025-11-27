package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupAssociatedClustersResponse Response Object
type UpdateClusterGroupAssociatedClustersResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateClusterGroupAssociatedClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupAssociatedClustersResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupAssociatedClustersResponse", string(data)}, " ")
}
