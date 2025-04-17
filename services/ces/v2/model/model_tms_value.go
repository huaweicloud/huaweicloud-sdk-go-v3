package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsValue TMS标签值规范。
type TmsValue struct {
}

func (o TmsValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsValue struct{}"
	}

	return strings.Join([]string{"TmsValue", string(data)}, " ")
}
