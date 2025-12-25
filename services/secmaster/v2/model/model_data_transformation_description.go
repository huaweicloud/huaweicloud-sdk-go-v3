package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataTransformationDescription 数据转换描述
type DataTransformationDescription struct {
}

func (o DataTransformationDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationDescription struct{}"
	}

	return strings.Join([]string{"DataTransformationDescription", string(data)}, " ")
}
