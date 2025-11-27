package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterRulesRequest Request Object
type UpdateClusterRulesRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`

	Body *UpdateAssociatedRulesRequest `json:"body,omitempty"`
}

func (o UpdateClusterRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterRulesRequest", string(data)}, " ")
}
