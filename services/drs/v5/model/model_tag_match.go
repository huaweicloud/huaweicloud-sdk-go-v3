package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagMatch 标签匹配信息体.
type TagMatch struct {

	// 键，例如键值对{“aaa”:\"bbb\"}的key为\"aaa\"
	Key string `json:"key"`

	// 值，例如键值对{“aaa”:\"bbb\"}的value为\"bbb\"
	Value string `json:"value"`
}

func (o TagMatch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMatch struct{}"
	}

	return strings.Join([]string{"TagMatch", string(data)}, " ")
}
