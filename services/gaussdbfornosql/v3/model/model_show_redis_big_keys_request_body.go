package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowRedisBigKeysRequestBody struct {

	// 索引位置偏移量，表示从查询到的大key列表偏移offset条数据后查询对应的大key信息。 取值大于或等于0。不传该参数时，查询偏移量默认为0，表示从第一条大key开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询个数上限值。 取值范围：1~100。不传该参数时，默认查询前100条大key。
	Limit *int32 `json:"limit,omitempty"`

	// 大key的类型,一个字符串列表,支持\"string\",\"hash\",\"zset\",\"set\",\"list\",\"stream\"六种类型。
	KeyTypes *[]string `json:"key_types,omitempty"`
}

func (o ShowRedisBigKeysRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisBigKeysRequestBody struct{}"
	}

	return strings.Join([]string{"ShowRedisBigKeysRequestBody", string(data)}, " ")
}
