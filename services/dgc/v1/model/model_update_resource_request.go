package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateResourceRequest struct {
	// 资源id.

	ResourceId string `json:"resource_id"`

	Body *ResourceInfo `json:"body,omitempty"`
}

func (o UpdateResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceRequest", string(data)}, " ")
}
