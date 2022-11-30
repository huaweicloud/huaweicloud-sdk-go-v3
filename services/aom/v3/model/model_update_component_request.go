package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateComponentRequest struct {

	// 组件id
	ComponentId string `json:"component_id"`

	Body *ComponentParam `json:"body,omitempty"`
}

func (o UpdateComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequest struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequest", string(data)}, " ")
}
