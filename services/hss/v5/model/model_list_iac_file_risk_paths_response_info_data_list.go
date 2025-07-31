package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIacFileRiskPathsResponseInfoDataList 文件风险路径信息
type ListIacFileRiskPathsResponseInfoDataList struct {

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源所属的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 风险命中的规则
	HitRule *string `json:"hit_rule,omitempty"`

	// 存在风险的路径
	HitPath *string `json:"hit_path,omitempty"`
}

func (o ListIacFileRiskPathsResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIacFileRiskPathsResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"ListIacFileRiskPathsResponseInfoDataList", string(data)}, " ")
}
