package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MappingJoinFieldVo struct {

	// 属性1ID，ID字符串。
	Field1Id string `json:"field1_id"`

	// 属性2ID，ID字符串。
	Field2Id string `json:"field2_id"`

	// 属性1名称。
	Field1Name string `json:"field1_name"`

	// 属性2名称。
	Field2Name string `json:"field2_name"`
}

func (o MappingJoinFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingJoinFieldVo struct{}"
	}

	return strings.Join([]string{"MappingJoinFieldVo", string(data)}, " ")
}
