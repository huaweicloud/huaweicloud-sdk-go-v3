package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 测试计划处理者信息
type TestPlanDetailOwner struct {
	// 测试计划处理者id

	Id *string `json:"id,omitempty"`
	// 测试计划处理者名称，优先返回nickName，不存在则返回userName

	Name *string `json:"name,omitempty"`
	// 测试计划处理者的昵称，当用户未设置昵称时不返回该字段

	NickName *string `json:"nick_name,omitempty"`
	// 测试计划处理者的用户名称

	UserName *string `json:"user_name,omitempty"`
}

func (o TestPlanDetailOwner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestPlanDetailOwner struct{}"
	}

	return strings.Join([]string{"TestPlanDetailOwner", string(data)}, " ")
}
