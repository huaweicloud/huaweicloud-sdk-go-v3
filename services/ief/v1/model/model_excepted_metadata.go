package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExceptedMetadata 期望值设置的时间信息
type ExceptedMetadata struct {

	// 属性类型标识，string|int|float|boolean（boolean类型为true或false），默认为string
	Type *string `json:"type,omitempty"`
}

func (o ExceptedMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceptedMetadata struct{}"
	}

	return strings.Join([]string{"ExceptedMetadata", string(data)}, " ")
}
