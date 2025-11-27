package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupAssociatedRulesRequest Request Object
type UpdateClusterGroupAssociatedRulesRequest struct {

	// 容器舰队ID
	Clustergroupid string `json:"clustergroupid"`

	Body *UpdateAssociatedRulesRequest `json:"body,omitempty"`
}

func (o UpdateClusterGroupAssociatedRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupAssociatedRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupAssociatedRulesRequest", string(data)}, " ")
}
