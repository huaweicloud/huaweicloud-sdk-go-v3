package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Error struct {

	// 错误详情
	Detail string `json:"detail"`

	// 铂金版实例ID，如果为空则表示是专业版实例。
	IefInstanceId string `json:"ief_instance_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 规则ID
	RuleId string `json:"rule_id"`

	// 错误发生的时间
	Time string `json:"time"`
}

func (o Error) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Error struct{}"
	}

	return strings.Join([]string{"Error", string(data)}, " ")
}
