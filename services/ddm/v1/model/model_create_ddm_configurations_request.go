package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmConfigurationsRequest Request Object
type CreateDdmConfigurationsRequest struct {
	Body *CreateConfigurationRequest `json:"body,omitempty"`
}

func (o CreateDdmConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"CreateDdmConfigurationsRequest", string(data)}, " ")
}
