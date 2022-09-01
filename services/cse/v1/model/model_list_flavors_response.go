package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {

	// 微服务引擎专享版规格总个数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 微服务引擎专享版规格详情
	Data           *[]FlavorBrief `json:"data,omitempty" xml:"data"`
	HttpStatusCode int            `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
