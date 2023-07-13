package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAppInstanceHistoryResponseDto 应用实例历史版本
type QueryAppInstanceHistoryResponseDto struct {

	// 应用ID
	AppId *string `json:"app_id,omitempty"`

	// 应用版本号
	AppVersion *string `json:"app_version,omitempty"`

	// 应用实例历史版本号
	Version *string `json:"version,omitempty"`

	// 应用实例chart配置
	Values *interface{} `json:"values,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o QueryAppInstanceHistoryResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAppInstanceHistoryResponseDto struct{}"
	}

	return strings.Join([]string{"QueryAppInstanceHistoryResponseDto", string(data)}, " ")
}
