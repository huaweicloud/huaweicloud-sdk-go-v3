package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Execution struct {

	// 记录ID
	Id int32 `json:"id"`

	// 策略ID
	PolicyId int32 `json:"policy_id"`

	// 状态
	Status string `json:"status"`

	// 状态详情
	StatusText string `json:"status_text"`

	// 总数
	Total int32 `json:"total"`

	// 失败数
	Failed int32 `json:"failed"`

	// 成功数
	Succeed int32 `json:"succeed"`

	// 进行中的数量
	InProgress int32 `json:"in_progress"`

	// 停止数
	Stopped int32 `json:"stopped"`

	// 触发类型
	Trigger string `json:"trigger"`

	// 创建时间
	CreatedAt string `json:"created_at"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (o Execution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Execution struct{}"
	}

	return strings.Join([]string{"Execution", string(data)}, " ")
}
