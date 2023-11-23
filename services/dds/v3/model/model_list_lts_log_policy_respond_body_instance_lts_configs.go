package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsLogPolicyRespondBodyInstanceLtsConfigs LTS日志配置信息只会出现已经配置过的日志类型。
type ListLtsLogPolicyRespondBodyInstanceLtsConfigs struct {
	Instance *ListLtsLogPolicyRespondBodyInstance `json:"instance,omitempty"`

	// LTS日志配置明细。从未设置LTS日志流，不返回该字段。
	LtsConfigs *[]ListLtsLogPolicyRespondBodyLtsConfigs `json:"lts_configs,omitempty"`
}

func (o ListLtsLogPolicyRespondBodyInstanceLtsConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsLogPolicyRespondBodyInstanceLtsConfigs struct{}"
	}

	return strings.Join([]string{"ListLtsLogPolicyRespondBodyInstanceLtsConfigs", string(data)}, " ")
}
