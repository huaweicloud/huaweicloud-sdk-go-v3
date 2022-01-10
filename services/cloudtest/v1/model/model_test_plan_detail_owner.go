package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划创建者信息
type TestPlanDetailOwner struct {
	// 测试计划创建者id

	Id *string `json:"id,omitempty"`
	// 测试计划创建者名称

	Name *string `json:"name,omitempty"`
}

func (o TestPlanDetailOwner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanDetailOwner struct{}"
	}

	return strings.Join([]string{"TestPlanDetailOwner", string(data)}, " ")
}
