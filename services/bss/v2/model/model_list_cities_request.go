package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCitiesRequest struct {
	// 语言。zh_CN：中文en_us：英文缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`
	// 省份编码。

	ProvinceCode string `json:"province_code"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的数量，最大1000。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCitiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCitiesRequest struct{}"
	}

	return strings.Join([]string{"ListCitiesRequest", string(data)}, " ")
}
