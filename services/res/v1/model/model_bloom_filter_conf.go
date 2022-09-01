package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BloomFilterConf struct {

	// 待过滤行为类型。
	Behaviors *[]string `json:"behaviors,omitempty" xml:"behaviors"`

	// 过滤时间。
	Interval *int32 `json:"interval,omitempty" xml:"interval"`
}

func (o BloomFilterConf) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BloomFilterConf struct{}"
	}

	return strings.Join([]string{"BloomFilterConf", string(data)}, " ")
}
