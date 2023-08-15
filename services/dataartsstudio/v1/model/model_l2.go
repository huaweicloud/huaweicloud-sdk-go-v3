package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// L2 主题域中文名
type L2 struct {
}

func (o L2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L2 struct{}"
	}

	return strings.Join([]string{"L2", string(data)}, " ")
}
