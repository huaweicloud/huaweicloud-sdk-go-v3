package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceBindingRequest Request Object
type UpdateResourceBindingRequest struct {

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *UpdateResourceBody `json:"body,omitempty"`
}

func (o UpdateResourceBindingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceBindingRequest struct{}"
	}

	return strings.Join([]string{"UpdateResourceBindingRequest", string(data)}, " ")
}
