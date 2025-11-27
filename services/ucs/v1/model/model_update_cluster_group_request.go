package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterGroupRequest Request Object
type UpdateClusterGroupRequest struct {

	// 容器舰队ID
	Clustergroupid string `json:"clustergroupid"`

	Body *UpdateClusterGroupDescriptionRequest `json:"body,omitempty"`
}

func (o UpdateClusterGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterGroupRequest", string(data)}, " ")
}
