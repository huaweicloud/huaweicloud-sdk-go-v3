package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGrantRequest Request Object
type UpdateGrantRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *GrantSecretReqBody `json:"body,omitempty"`
}

func (o UpdateGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGrantRequest struct{}"
	}

	return strings.Join([]string{"UpdateGrantRequest", string(data)}, " ")
}
