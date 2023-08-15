package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BloomFilterConf
type BloomFilterConf struct {

	// 待过滤行为类型。
	Behaviors *[]string `json:"behaviors,omitempty"`

	// 过滤时间。
	Interval *int32 `json:"interval,omitempty"`
}

func (o BloomFilterConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BloomFilterConf struct{}"
	}

	return strings.Join([]string{"BloomFilterConf", string(data)}, " ")
}
