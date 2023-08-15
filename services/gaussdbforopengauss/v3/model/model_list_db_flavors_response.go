package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbFlavorsResponse Response Object
type ListDbFlavorsResponse struct {

	// 实例规格信息。
	Flavors *[]FlavorResult `json:"flavors,omitempty"`

	// 查询总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListDbFlavorsResponse", string(data)}, " ")
}
