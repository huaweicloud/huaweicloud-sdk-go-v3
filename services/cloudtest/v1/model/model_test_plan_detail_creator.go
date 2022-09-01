package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划创建者信息
type TestPlanDetailCreator struct {

	// 测试计划创建者id
	Id *string `json:"id,omitempty" xml:"id"`

	// 测试计划创建者的昵称，当用户未设置昵称时不返回该字段
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	// 测试计划创建者的用户名称
	UserName *string `json:"user_name,omitempty" xml:"user_name"`
}

func (o TestPlanDetailCreator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanDetailCreator struct{}"
	}

	return strings.Join([]string{"TestPlanDetailCreator", string(data)}, " ")
}
