package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataClassRefPojo 数据类对象信息
type DataClassRefPojo struct {

	// 数据类ID
	Id string `json:"id"`

	// 数据类名称
	Name *string `json:"name,omitempty"`
}

func (o DataClassRefPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassRefPojo struct{}"
	}

	return strings.Join([]string{"DataClassRefPojo", string(data)}, " ")
}
