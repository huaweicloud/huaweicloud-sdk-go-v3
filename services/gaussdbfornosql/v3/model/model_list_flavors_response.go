package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 实例规格信息列表。
	Flavors        *[]ListFlavorsResult `json:"flavors,omitempty" xml:"flavors"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
