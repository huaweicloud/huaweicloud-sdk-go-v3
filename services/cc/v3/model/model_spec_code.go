package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecCode 带宽包产品规格编码。
type SpecCode struct {

	// 带宽包实例的规格编码。
	SpecCode string `json:"spec_code"`
}

func (o SpecCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecCode struct{}"
	}

	return strings.Join([]string{"SpecCode", string(data)}, " ")
}
