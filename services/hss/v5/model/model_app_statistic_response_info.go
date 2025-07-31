package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppStatisticResponseInfo 进程统计信息
type AppStatisticResponseInfo struct {

	// **参数解释**: 软件名称 **取值范围**: 字符长度1-128位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 进程数量 **取值范围**: 字符长度0-100000位
	Num *int32 `json:"num,omitempty"`
}

func (o AppStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"AppStatisticResponseInfo", string(data)}, " ")
}
