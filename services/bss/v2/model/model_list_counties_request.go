package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCountiesRequest struct {
	// 语言。zh_CN：中文en_us：英文缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`
	// 城市的编码。

	CityCode string `json:"city_code"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的数量，最大1000。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCountiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCountiesRequest struct{}"
	}

	return strings.Join([]string{"ListCountiesRequest", string(data)}, " ")
}
