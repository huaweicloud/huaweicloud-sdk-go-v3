package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 静态属性的元数据信息，默认为{\"type\": \"string\"}
type ValueInPropertyVisitorsRegisterTypeMetadata struct {
	// 属性类型标识，string|int|float|boolean（boolean类型为true或false），默认为string

	Type string `json:"type"`
}

func (o ValueInPropertyVisitorsRegisterTypeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValueInPropertyVisitorsRegisterTypeMetadata struct{}"
	}

	return strings.Join([]string{"ValueInPropertyVisitorsRegisterTypeMetadata", string(data)}, " ")
}
