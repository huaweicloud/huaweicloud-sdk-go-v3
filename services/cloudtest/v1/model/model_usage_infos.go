package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UsageInfos 资源已用容量对象
type UsageInfos struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 资源标识
	Id *string `json:"id,omitempty"`

	// 资源总量
	Amount *string `json:"amount,omitempty"`

	// 已消耗用量
	Used *string `json:"used,omitempty"`

	// 资源已用容量百分比,例如80% 值为80
	UsedPercent *int32 `json:"used_percent,omitempty"`

	// 版本超限信息
	UsageInfo *[]UsageInfos `json:"usage_info,omitempty"`
}

func (o UsageInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageInfos struct{}"
	}

	return strings.Join([]string{"UsageInfos", string(data)}, " ")
}
