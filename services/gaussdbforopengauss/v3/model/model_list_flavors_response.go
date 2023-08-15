package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsResponse Response Object
type ListFlavorsResponse struct {

	// 实例规格信息。
	Flavors *[]Flavor `json:"flavors,omitempty"`

	// 查询的记录总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
