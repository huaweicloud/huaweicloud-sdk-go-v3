package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Curve struct {

	// 入带宽
	In *float32 `json:"in,omitempty"`

	// 出带宽
	Out *float32 `json:"out,omitempty"`

	// 时间戳
	Time *int32 `json:"time,omitempty"`
}

func (o Curve) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Curve struct{}"
	}

	return strings.Join([]string{"Curve", string(data)}, " ")
}
