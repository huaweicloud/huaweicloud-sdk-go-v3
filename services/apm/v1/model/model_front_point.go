package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据点模型
type FrontPoint struct {

	// 时间
	Time *int64 `json:"time,omitempty" xml:"time"`

	// 值
	Value *interface{} `json:"value,omitempty" xml:"value"`
}

func (o FrontPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontPoint struct{}"
	}

	return strings.Join([]string{"FrontPoint", string(data)}, " ")
}
