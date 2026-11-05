package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiApplicationListItem 应用列表项（精简）。
type AiApplicationListItem struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	// 关联的PG实例ID。
	ManagedInstanceId *string `json:"managed_instance_id,omitempty"`
}

func (o AiApplicationListItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiApplicationListItem struct{}"
	}

	return strings.Join([]string{"AiApplicationListItem", string(data)}, " ")
}
