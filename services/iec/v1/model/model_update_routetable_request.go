package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRoutetableRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id" xml:"routetable_id"`

	Body *UpdateRoutetableRequesBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateRoutetableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutetableRequest struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableRequest", string(data)}, " ")
}
