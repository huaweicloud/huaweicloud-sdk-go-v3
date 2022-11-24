package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRoutetablesResponse struct {

	// 路由表
	Routetables *[]ListRoutetableOption `json:"routetables,omitempty"`

	// 数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRoutetablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRoutetablesResponse struct{}"
	}

	return strings.Join([]string{"ListRoutetablesResponse", string(data)}, " ")
}
