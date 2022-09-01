package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源历史
type HistoryItem struct {

	// 用户id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type"`

	// 该资源在RMS系统捕获时间
	CaptureTime *string `json:"capture_time,omitempty" xml:"capture_time"`

	// 资源状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 资源关系列表
	Relations *[]ResourceRelation `json:"relations,omitempty" xml:"relations"`

	Resource *ResourceEntity `json:"resource,omitempty" xml:"resource"`
}

func (o HistoryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryItem struct{}"
	}

	return strings.Join([]string{"HistoryItem", string(data)}, " ")
}
