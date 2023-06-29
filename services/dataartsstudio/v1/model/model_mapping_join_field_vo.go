package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MappingJoinFieldVo struct {

	// 属性id
	Field1Id int64 `json:"field1_id"`

	// 属性id
	Field2Id int64 `json:"field2_id"`

	// 名称
	Field1Name string `json:"field1_name"`

	// 名称
	Field2Name string `json:"field2_name"`
}

func (o MappingJoinFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingJoinFieldVo struct{}"
	}

	return strings.Join([]string{"MappingJoinFieldVo", string(data)}, " ")
}
