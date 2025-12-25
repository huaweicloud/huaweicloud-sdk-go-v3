package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataTransformationName 数据转换名称
type DataTransformationName struct {
}

func (o DataTransformationName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationName struct{}"
	}

	return strings.Join([]string{"DataTransformationName", string(data)}, " ")
}
