package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRoutetableRequest Request Object
type DeleteRoutetableRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id"`
}

func (o DeleteRoutetableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRoutetableRequest struct{}"
	}

	return strings.Join([]string{"DeleteRoutetableRequest", string(data)}, " ")
}
