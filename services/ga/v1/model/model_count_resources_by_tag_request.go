package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesByTagRequest Request Object
type CountResourcesByTagRequest struct {

	// 资源类型。
	ResourceType *ResourceType `json:"resource_type"`

	Body *ListResourcesByTagRequestBody `json:"body,omitempty"`
}

func (o CountResourcesByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesByTagRequest struct{}"
	}

	return strings.Join([]string{"CountResourcesByTagRequest", string(data)}, " ")
}
