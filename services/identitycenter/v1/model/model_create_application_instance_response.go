package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationInstanceResponse Response Object
type CreateApplicationInstanceResponse struct {
	ApplicationInstance *ApplicationInstanceDto `json:"application_instance,omitempty"`
	HttpStatusCode      int                     `json:"-"`
}

func (o CreateApplicationInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationInstanceResponse", string(data)}, " ")
}
