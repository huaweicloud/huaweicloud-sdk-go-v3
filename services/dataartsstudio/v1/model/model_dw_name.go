package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DwName 数据连接名称
type DwName struct {
}

func (o DwName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DwName struct{}"
	}

	return strings.Join([]string{"DwName", string(data)}, " ")
}
