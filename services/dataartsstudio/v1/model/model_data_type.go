package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataType 字段类型。
type DataType struct {
}

func (o DataType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataType struct{}"
	}

	return strings.Join([]string{"DataType", string(data)}, " ")
}
