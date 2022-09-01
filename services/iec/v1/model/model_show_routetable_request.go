package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRoutetableRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id" xml:"routetable_id"`
}

func (o ShowRoutetableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRoutetableRequest struct{}"
	}

	return strings.Join([]string{"ShowRoutetableRequest", string(data)}, " ")
}
