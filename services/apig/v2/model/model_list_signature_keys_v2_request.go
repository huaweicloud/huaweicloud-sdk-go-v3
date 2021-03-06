package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSignatureKeysV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 签名密钥编号

	Id *string `json:"id,omitempty"`
	// 签名密钥名称

	Name *string `json:"name,omitempty"`
	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0

	Offset *int64 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
	// 指定需要精确匹配查找的参数名称，目前仅支持name

	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListSignatureKeysV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSignatureKeysV2Request struct{}"
	}

	return strings.Join([]string{"ListSignatureKeysV2Request", string(data)}, " ")
}
