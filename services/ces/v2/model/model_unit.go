package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Unit 数据的单位。
type Unit struct {
}

func (o Unit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Unit struct{}"
	}

	return strings.Join([]string{"Unit", string(data)}, " ")
}
