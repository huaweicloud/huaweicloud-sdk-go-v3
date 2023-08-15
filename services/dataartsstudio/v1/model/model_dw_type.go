package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DwType 数据连接类型
type DwType struct {
}

func (o DwType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DwType struct{}"
	}

	return strings.Join([]string{"DwType", string(data)}, " ")
}
