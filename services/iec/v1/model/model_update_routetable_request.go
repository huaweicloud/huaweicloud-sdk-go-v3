package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoutetableRequest Request Object
type UpdateRoutetableRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id"`

	Body *UpdateRoutetableRequesBody `json:"body,omitempty"`
}

func (o UpdateRoutetableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutetableRequest struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableRequest", string(data)}, " ")
}
