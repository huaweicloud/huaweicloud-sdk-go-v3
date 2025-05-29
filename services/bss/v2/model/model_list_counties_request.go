package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCountiesRequest Request Object
type ListCountiesRequest struct {

	// 语言，字段预留,目前仅支持中文。默认zh_cn，枚举：zh_cn：中文 en_us：英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 城市的编码。
	CityCode string `json:"city_code"`

	// 偏移量，从0开始。此参数不携带或携带值为空或携带值为null时，取默认值0；不支持携带值为空串。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量，最大1000。该参数不携带时，取默认值10。该参数不支持携带值为空、不支持携带值为空串、不支持携带值为null。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCountiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCountiesRequest struct{}"
	}

	return strings.Join([]string{"ListCountiesRequest", string(data)}, " ")
}
