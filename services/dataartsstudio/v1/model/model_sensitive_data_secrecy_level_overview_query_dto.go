package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SensitiveDataSecrecyLevelOverviewQueryDto 敏感规则数据结果表查询对象
type SensitiveDataSecrecyLevelOverviewQueryDto struct {

	// 密级ID
	SecrecyLevelId *string `json:"secrecy_level_id,omitempty"`

	// 密级名称
	SecrecyLevelName *string `json:"secrecy_level_name,omitempty"`

	// 密级的等级
	SecrecyLevelNumber *int64 `json:"secrecy_level_number,omitempty"`

	// 当前密级下的敏感字段数量
	Count *int64 `json:"count,omitempty"`
}

func (o SensitiveDataSecrecyLevelOverviewQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SensitiveDataSecrecyLevelOverviewQueryDto struct{}"
	}

	return strings.Join([]string{"SensitiveDataSecrecyLevelOverviewQueryDto", string(data)}, " ")
}
