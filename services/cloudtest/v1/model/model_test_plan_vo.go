package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestPlanVo 测试计划信息
type TestPlanVo struct {

	// 测试计划URI
	Uri *string `json:"uri,omitempty"`

	// 测试计划名称
	Name *string `json:"name,omitempty"`
}

func (o TestPlanVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanVo struct{}"
	}

	return strings.Join([]string{"TestPlanVo", string(data)}, " ")
}
