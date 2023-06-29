package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataClassRefPojo 数据类对象信息
type DataClassRefPojo struct {

	// 唯一标识ID
	Id string `json:"id"`

	// 唯一标识ID
	Name *string `json:"name,omitempty"`
}

func (o DataClassRefPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassRefPojo struct{}"
	}

	return strings.Join([]string{"DataClassRefPojo", string(data)}, " ")
}
