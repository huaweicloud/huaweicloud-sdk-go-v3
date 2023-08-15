package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrontPoint 数据点模型。
type FrontPoint struct {

	// 时间。
	Time *int64 `json:"time,omitempty"`

	// 值。
	Value *interface{} `json:"value,omitempty"`
}

func (o FrontPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrontPoint struct{}"
	}

	return strings.Join([]string{"FrontPoint", string(data)}, " ")
}
