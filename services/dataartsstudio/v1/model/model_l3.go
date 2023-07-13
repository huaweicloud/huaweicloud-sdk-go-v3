package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// L3 业务对象中文名
type L3 struct {
}

func (o L3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L3 struct{}"
	}

	return strings.Join([]string{"L3", string(data)}, " ")
}
