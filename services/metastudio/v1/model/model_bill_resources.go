package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BillResources 计费资源信息
type BillResources struct {
	ResourceType *BillResourceType `json:"resource_type,omitempty"`

	// 资源数量。
	ResourceNums *int32 `json:"resource_nums,omitempty"`
}

func (o BillResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BillResources struct{}"
	}

	return strings.Join([]string{"BillResources", string(data)}, " ")
}
