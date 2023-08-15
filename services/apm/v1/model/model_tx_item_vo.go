package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TxItemVo URL跟踪视图。
type TxItemVo struct {

	// 组件名称。
	AppName *string `json:"app_name,omitempty"`

	// 环境名称。
	EnvName *string `json:"env_name,omitempty"`

	// 事务显示名称。
	TxDisplayName *string `json:"tx_display_name,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 组件id。
	AppId *int64 `json:"app_id,omitempty"`

	// 事务名称。
	TxName *string `json:"tx_name,omitempty"`

	// 调用次数。
	InvokeCount *int32 `json:"invoke_count,omitempty"`

	// 总耗时。
	TotalTime *int32 `json:"total_time,omitempty"`

	// 错误次数。
	ErrorCount *int32 `json:"error_count,omitempty"`
}

func (o TxItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TxItemVo struct{}"
	}

	return strings.Join([]string{"TxItemVo", string(data)}, " ")
}
