package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceSearchParam 获取实例信息列表入参。
type InstanceSearchParam struct {

	// 环境id。
	EnvId int64 `json:"env_id"`

	// 当前页码。
	Page int32 `json:"page"`

	// 每页数据容量。
	PageSize *int32 `json:"page_size,omitempty"`

	// 关键字。
	Keyword *string `json:"keyword,omitempty"`

	// 实例状态。
	Status *int32 `json:"status,omitempty"`

	// 是否返回计数结果。
	ReturnCount *bool `json:"return_count,omitempty"`
}

func (o InstanceSearchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSearchParam struct{}"
	}

	return strings.Join([]string{"InstanceSearchParam", string(data)}, " ")
}
