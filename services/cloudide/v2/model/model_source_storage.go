package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SourceStorage struct {

	// 位置
	Location *string `json:"location,omitempty" xml:"location"`

	// 参数值
	Parameters map[string]string `json:"parameters,omitempty" xml:"parameters"`

	// 类型
	Type *string `json:"type,omitempty" xml:"type"`
}

func (o SourceStorage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceStorage struct{}"
	}

	return strings.Join([]string{"SourceStorage", string(data)}, " ")
}
