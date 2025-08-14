package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetApplicationInstanceResponse Response Object
type GetApplicationInstanceResponse struct {
	ApplicationInstance *ApplicationInstanceDto `json:"application_instance,omitempty"`
	HttpStatusCode      int                     `json:"-"`
}

func (o GetApplicationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetApplicationInstanceResponse struct{}"
	}

	return strings.Join([]string{"GetApplicationInstanceResponse", string(data)}, " ")
}
