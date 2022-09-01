package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateExtensionAuthorizationRequest struct {

	// CloudIDE实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ExtensionAuthorization `json:"body,omitempty" xml:"body"`
}

func (o CreateExtensionAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExtensionAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"CreateExtensionAuthorizationRequest", string(data)}, " ")
}
