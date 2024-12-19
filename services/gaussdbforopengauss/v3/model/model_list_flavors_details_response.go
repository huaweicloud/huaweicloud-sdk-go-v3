package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsDetailsResponse Response Object
type ListFlavorsDetailsResponse struct {

	// 实例规格信息。
	Flavors *[]FlavorResult `json:"flavors,omitempty"`

	// 查询总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFlavorsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsDetailsResponse", string(data)}, " ")
}
