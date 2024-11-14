package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryDbCacheMappingResponse 内存加速映射信息。
type QueryDbCacheMappingResponse struct {

	// 内存加速映射ID。
	Id string `json:"id"`

	// 内存加速映射名称。
	Name string `json:"name"`

	// 源实例ID。
	SourceInstanceId string `json:"source_instance_id"`

	// 源实例名称。
	SourceInstanceName string `json:"source_instance_name"`

	// 目标实例ID。
	TargetInstanceId string `json:"target_instance_id"`

	// 目标实例名称。
	TargetInstanceName string `json:"target_instance_name"`

	// 内存加速映射关系。  - normal： 表示内存加速映射正常。  - creating： 表示内存加速映射创建中。  - createfail： 表示内存加速映射创建失败。  - deleting： 表示内存加速映射解除中。  - stopped： 表示内存加速映射停止。  - deleted： 表示内存加速映射已解除。
	Status string `json:"status"`

	// 内存加速映射创建时间。
	Created string `json:"created"`

	// 内存加速映射最新变更的时间。
	Updated string `json:"updated"`

	// 该内存加速映射下的规则数量。
	RuleCount *int32 `json:"rule_count,omitempty"`
}

func (o QueryDbCacheMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDbCacheMappingResponse struct{}"
	}

	return strings.Join([]string{"QueryDbCacheMappingResponse", string(data)}, " ")
}
