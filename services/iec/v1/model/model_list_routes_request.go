package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRoutesRequest struct {

	// 路由表ID
	RoutetableId string `json:"routetable_id"`
}

func (o ListRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutesRequest struct{}"
	}

	return strings.Join([]string{"ListRoutesRequest", string(data)}, " ")
}
