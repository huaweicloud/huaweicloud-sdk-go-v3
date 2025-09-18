package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GcbBorderCross 全域互联带宽跨境属性。
type GcbBorderCross struct {

	// 全域互联带宽跨境属性。
	Bordercross bool `json:"bordercross"`
}

func (o GcbBorderCross) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbBorderCross struct{}"
	}

	return strings.Join([]string{"GcbBorderCross", string(data)}, " ")
}
