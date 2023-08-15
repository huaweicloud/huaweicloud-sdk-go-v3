package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Metadata struct {

	// 属性类型标识，string|int|float|boolean（boolean类型为true或false），默认为string
	Type *string `json:"type,omitempty"`
}

func (o Metadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Metadata struct{}"
	}

	return strings.Join([]string{"Metadata", string(data)}, " ")
}
