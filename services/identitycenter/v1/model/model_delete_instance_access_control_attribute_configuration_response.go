package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceAccessControlAttributeConfigurationResponse Response Object
type DeleteInstanceAccessControlAttributeConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceAccessControlAttributeConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceAccessControlAttributeConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceAccessControlAttributeConfigurationResponse", string(data)}, " ")
}
