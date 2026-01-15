package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDomainControllerRequest Request Object
type ValidateDomainControllerRequest struct {
	Body *ValidateDcRequestBody `json:"body,omitempty"`
}

func (o ValidateDomainControllerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDomainControllerRequest struct{}"
	}

	return strings.Join([]string{"ValidateDomainControllerRequest", string(data)}, " ")
}
