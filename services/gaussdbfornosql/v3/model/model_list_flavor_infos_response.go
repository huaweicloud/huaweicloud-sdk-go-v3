package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorInfosResponse Response Object
type ListFlavorInfosResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 实例规格信息列表。
	Flavors        *[]ListFlavorsResult `json:"flavors,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListFlavorInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorInfosResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorInfosResponse", string(data)}, " ")
}
