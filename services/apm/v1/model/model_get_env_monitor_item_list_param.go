package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询特定环境下监控项参数模型
type GetEnvMonitorItemListParam struct {

	// 环境id
	EnvId int64 `json:"env_id" xml:"env_id"`

	// 页码
	Page int32 `json:"page" xml:"page"`

	// 每页数量
	PageSize int32 `json:"page_size" xml:"page_size"`
}

func (o GetEnvMonitorItemListParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetEnvMonitorItemListParam struct{}"
	}

	return strings.Join([]string{"GetEnvMonitorItemListParam", string(data)}, " ")
}
