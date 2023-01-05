package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSearchCriteriasBody struct {

	// 快速查询字段
	Criteria string `json:"criteria"`

	// 企业项目id
	EpsId *string `json:"eps_id,omitempty"`

	// 创建快速查询名称
	Name string `json:"name"`

	// 查询类型 原始日志：ORIGINALLOG 可视化日志: VISUALIZATION
	SearchType string `json:"search_type"`
}

func (o CreateSearchCriteriasBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchCriteriasBody struct{}"
	}

	return strings.Join([]string{"CreateSearchCriteriasBody", string(data)}, " ")
}
