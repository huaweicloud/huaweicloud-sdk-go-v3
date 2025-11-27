package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupAssociatedClustersRequest Request Object
type UpdateClusterGroupAssociatedClustersRequest struct {

	// 容器舰队ID
	Clustergroupid string `json:"clustergroupid"`

	Body *UpdateClusterGroupAssociatedClustersRequestBody `json:"body,omitempty"`
}

func (o UpdateClusterGroupAssociatedClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupAssociatedClustersRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupAssociatedClustersRequest", string(data)}, " ")
}
