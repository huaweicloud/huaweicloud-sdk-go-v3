package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataclassName 数据类名称
type DataclassName struct {
}

func (o DataclassName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataclassName struct{}"
	}

	return strings.Join([]string{"DataclassName", string(data)}, " ")
}
