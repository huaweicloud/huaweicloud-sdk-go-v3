package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataclassInfoRef 数据类信息
type DataclassInfoRef struct {

	// 数据类ID
	Id *string `json:"id,omitempty"`

	// 数据类名称
	Name *string `json:"name,omitempty"`
}

func (o DataclassInfoRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataclassInfoRef struct{}"
	}

	return strings.Join([]string{"DataclassInfoRef", string(data)}, " ")
}
