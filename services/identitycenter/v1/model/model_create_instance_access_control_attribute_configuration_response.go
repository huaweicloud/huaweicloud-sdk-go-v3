package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceAccessControlAttributeConfigurationResponse Response Object
type CreateInstanceAccessControlAttributeConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateInstanceAccessControlAttributeConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceAccessControlAttributeConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceAccessControlAttributeConfigurationResponse", string(data)}, " ")
}
