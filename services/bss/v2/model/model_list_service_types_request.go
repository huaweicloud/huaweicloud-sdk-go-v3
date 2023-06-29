package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceTypesRequest Request Object
type ListServiceTypesRequest struct {

	// 语言。zh_CN：中文en_US：英文缺省为zh_CN。
	XLanguage *string `json:"X-Language,omitempty"`

	// 每次查询的数量，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListServiceTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceTypesRequest struct{}"
	}

	return strings.Join([]string{"ListServiceTypesRequest", string(data)}, " ")
}
