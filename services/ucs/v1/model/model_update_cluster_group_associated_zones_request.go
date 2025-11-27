package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupAssociatedZonesRequest Request Object
type UpdateClusterGroupAssociatedZonesRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	Body *UpdateAssociatedZonesRequest `json:"body,omitempty"`
}

func (o UpdateClusterGroupAssociatedZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupAssociatedZonesRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupAssociatedZonesRequest", string(data)}, " ")
}
