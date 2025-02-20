package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceMultiCountResponseData struct {

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源数量
	Count *int32 `json:"count,omitempty"`
}

func (o ResourceMultiCountResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceMultiCountResponseData struct{}"
	}

	return strings.Join([]string{"ResourceMultiCountResponseData", string(data)}, " ")
}
