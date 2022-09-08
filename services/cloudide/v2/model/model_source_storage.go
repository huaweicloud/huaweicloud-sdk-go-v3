package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SourceStorage struct {

	// 位置
	Location *string `json:"location,omitempty"`

	// 参数值
	Parameters map[string]string `json:"parameters,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`
}

func (o SourceStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceStorage struct{}"
	}

	return strings.Join([]string{"SourceStorage", string(data)}, " ")
}
